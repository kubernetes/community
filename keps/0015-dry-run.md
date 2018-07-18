---
kep-number: 15
title: Dry-run
authors:
  - "@apelisse"
owning-sig: sig-api-machinery
participating-sigs:
  - sig-api-machinery
  - sig-cli
reviewers:
  - "@lavalamp"
  - "@deads2k"
approvers:
  - "@erictune"
editor: apelisse
creation-date: 2018-06-21
last-updated: 2018-06-21
status: provisional
---
# Kubernetes Dry-run

Dry-run is a new feature that we intend to implement in the api-server. The goal
is to be able to send requests to modifying endpoints, and see if the request
would have succeeded (admission chain, validation, merge conflicts, ...) and/or
what would have happened without having it actually happen. The response body
for the request should be as close as possible to a non dry-run response.

## Specifying dry-run

Dry-run is triggered by setting the “dryRun” query parameter on modifying
verbs: POST, PUT, PATCH and DELETE.

This parameter is a string, working as an enum:
- All: Everything will run as normal, except for the storage that won’t be
  modified. Everything else should work as expected: admission controllers will
  be run to check that the request is valid, mutating controllers will change
  the object, merge will be performed on PATCH. The storage layer will be
  informed not to save, and the final object will be returned to the user with
  normal status code.
- Leave the value empty, or don't specify the parameter at all to keep the
  default modifying behavior.

No other values are supported yet, but this gives us an opportunity to create a
finer-grained mechanism later, if necessary.

## Admission controllers

Admission controllers need to be modified to understand that the request is a
“dry-run” request. While admission controllers are not supposed to have
side-effects when triggered, some of them still do. Quotas for example uses the
current request values to change the available quotas. Providing a parameter,
either as an argument for built-in admission plugins or through a dryRun
query-parameter for dynamic webhooks, will give them a chance not to have any
side-effect.

All admission controllers will have to be verified and changed. A new flag will
be added so that webhooks can explicitly register with dry-run support. If
dry-run is requested on a non-supported webhook, the request will be completely
rejected.

## Storage

The storage layer will be modified, so that it can know if request is dry-run,
most likely by looking for the field in the “Options” structure (missing for
some handlers, to be added). If it is, it will NOT store the object, but return
success. That success can be forwarded back to the user.
