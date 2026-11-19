# SIG Slack lead generator and verifier

This tool keeps the SIG lead Slack usergroups in
`communication/slack-config/sigslackleads.yaml` aligned with the active
leadership in `sigs.yaml`. It is offline and never changes live Slack state.

Run `go run ./generator/sigslackleads/cmd` to verify without writing. To
reconcile the generated file, run `go run ./generator/sigslackleads/cmd
-generate`.
`make generate` performs this generation after the other `sigs.yaml`-derived
outputs. The verification script remains read-only and is included by
`make verify`.

## Generated usergroups

Every usergroup in `sigslackleads.yaml` is owned by the generator. The default
handle is `<SIG dir>-leads`. A SIG can override the generated handle through the
optional `contact.slack_usergroup` field in `sigs.yaml`:

```yaml
contact:
  slack: sig-example
  slack_usergroup: example-leads
```

A SIG can provide a custom description through
`contact.slack_usergroup_description` when the default generated text is not
appropriate. It can also require additional auto-join channels through
`contact.slack_usergroup_channels`:

```yaml
contact:
  slack: sig-example
  slack_usergroup_description: Leads for SIG Example. Ping for Example questions.
  slack_usergroup_channels:
    - example-alerts
```

`contact.slack` is always required and is always included in the generated
usergroup channels. A generated handle that collides with a channel fails
validation until `contact.slack_usergroup` selects a non-conflicting handle.

## Verification rules

The verifier:

- derives active members from `leadership.chairs` and `leadership.tech_leads`;
- excludes emeritus leadership;
- resolves every GitHub login to exactly one case-insensitive `users.yaml` key;
- requires valid, distinct Slack IDs for active leaders;
- requires the sorted, deduplicated member list to match exactly;
- requires every configured group channel, including SIG-specific additions;
- rejects missing, extra, duplicate, external, or channel-colliding groups;
- rejects generated handles that duplicate unrelated usergroups.

Leadership, handle, and channel changes should be made in `sigs.yaml`, followed
by `make generate`. Add any new confirmed GitHub-login to Slack-ID mapping in
`users.yaml` before generation.
