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
status: implementable
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
“dry-run” request. Admission controllers are allowed to have side-effects
when triggered, as long as there is a reconciliation system, because it is not
guaranteed that subsequent validating will permit the request to finish.
Quotas for example uses the current request values to change the available quotas.
The ```admission.Attributes``` interface will be edited like this, to inform the
built-in admission controllers if a request is a dry-run:
```golang
type Attributes interface {
	...
	// IsDryRun indicates that modifications will definitely not be persisted for this request. This is to prevent
	// admission controllers with side effects and a method of reconciliation from being overwhelmed.
	// However, a value of false for this does not mean that the modification will be persisted, because it
	// could still be rejected by a subsequent validation step.
	IsDryRun() bool
	...
}
```

All built-in admission controllers will then have to be checked, and the ones with side
effects will have to be changed to handle the dry-run case correctly. Some examples of
built-in admission controllers with the possibility for side-effects are:
- ResourceQuota
- EventRateLimit
- NamespaceAutoProvision
- (Valid|Mut)atingAdmissionWebhook

To address the possibility of webhook authors [relying on side effects](https://github.com/kubernetes/website/blame/836629cb118e0f74545cc7d6d97aa6b9edfa1a16/content/en/docs/reference/access-authn-authz/admission-controllers.md#L582-L584), a new field
will be added to ```admissionregistration.k8s.io/v1beta1.ValidatingWebhookConfiguration``` and
```admissionregistration.k8s.io/v1beta1.MutatingWebhookConfiguration``` so that webhooks
can explicitly register as having dry-run support.
If dry-run is requested on a non-supported webhook, the request will be completely rejected,
as a 400: Bad Request. This field will be defaulted to true and deprecated in v1, and completely removed in v2.
All webhooks registered with v2 will be assumed to support dry run. The [api conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md) advise
against bool fields because "many ideas start as boolean but eventually trend towards a small set
of mutually exclusive options" but in this case, we plan to remove the field in a future version.
```golang
// admissionregistration.k8s.io/v1beta1
...
type Webhook struct {
	...
	// DryRunnable defines whether this webhook will correctly handle dryRun requests.
	// If false, any dryRun requests to resources/subresources this webhook applies to
	// will be completely rejected and the webhook will not be called.
	// Defaults to false.
	// +optional
	DryRunnable *bool `json:"dryRunnable,omitempty" protobuf:"varint,6,number,opt,name=dryRunnable"`
}
```

Additionally, a new field will be added to ```admission.k8s.io/v1beta1.AdmissionReview```
API object to reflect the changes to the ```admission.Attributes``` interface, indicating
whether or not the request being reviewed is for a dry-run:
```golang
// admission.k8s.io/v1beta1
...
type AdmissionRequest struct {
	...
	// DryRun indicates that modifications will definitely not be persisted for this request.
	// Defaults to false.
	// +optional
	DryRun *bool `json:"dryRun,omitempty" protobuf:"varint,11,number,opt,name=dryRun"`
}
```

## Generated values

Some values of the object are typically generated before the object is persisted:
- generateName can be used to assign a unique random name to the object,
- creationTimestamp/deletionTimestamp records the time of creation/deletion,
- UID uniquely identifies the object and is randomly generated (non-deterministic),
- resourceVersion tracks the persisted version of the object.

Most of these values are not useful in the context of dry-run, and could create
some confusion. The UID and the generated name would have a different value in a
dry-run and non-dry-run creation. These values will be left empty when
performing a dry-run.

CreationTimestamp and DeletionTimestamp are also generated on creation/deletion,
but there are less ways to abuse them so they will be generated as they for a
regular request.

ResourceVersion will also be left empty on creation. On updates, the value will
stay unchanged.

## Storage

The storage layer will be modified, so that it can know if request is dry-run,
most likely by looking for the field in the “Options” structure (missing for
some handlers, to be added). If it is, it will NOT store the object, but return
success. That success can be forwarded back to the user.

A dry-run request should behave as close as possible to a regular
request. Attempting to dry-run create an existing object will result in an
`AlreadyExists` error to be returned. Similarly, if a dry-run update is
performed on a non-existing object, a `NotFound` error will be returned.
