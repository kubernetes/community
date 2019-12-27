#### Presenter:James Munnelly

#### Topic: Cert-Manager

#### Date & Time: 11:30

#### Notes

Are there different ways that Kubernetes consumes secrets? Because there are different ways Kubernetes uses secrets.

With webhooks and ca-bundle needing to be inline is approached it ca-injector.

You add annotation into webhook resource and inject ca-crt for reference.
Cert manager uses it for its own TLS injections.

People are giving their permissions to the daemonsets which have ability to query all secrets anyways, so there should
be some separating at the daemonset level to restrict access to secrets across the cluster.

Like the node emission controller.

The CSR api should be extended, there has been interest in adding certificate class.
Originally there was just a certificate resource which represented long running cert.

Looking at adding approval controllers

Put together CSI driver for that specific request api which you can add volume to your pod which will go off in the
background request certificate api for new certificate. Using this you could take the pod identity to make the request,
and a approval system that was watching the pods and approve it when it sees request

Security issues are still a problem with certmanager because of the access to api secrets and the lack of restriction
to individual secrets

Node authorizer + node restriction plugin this will help restrict secrets by pod the driver uses the pods identity to 
create the request.

There is a csi driver for reading vault secret and inserting them into your pod, uses its own driver and you have to
grant it all permissions.

With service ca approach with open shift you just create annotations and it will handle everything for you

That is also what csidriver is doing you say you want a secret and then you just get one.

Annotating a service could leverage a webhook to get a certificate with that services name, at the service level it
would make CA and then you would get the certificates. Per pod basics and you would throw the CA in the ingress.

How do people see themselves using csi driver

Since 1.12 service accounts have been injected but it doesn't solve service authentication
people tend to use Kubernetes as auth mechanism for service to service auth but not necessarily is the best method
should creat infrastructure to do that properly.

Control the setup of who or what service accounts issues these requests. Vault doesn't behave like a kubernetes api and
this is why it would be a better choice

Need a way to control certificate requests and permissions across services and ingress to not give full access and
restrict by pod rather than namespace???

CSI drivers assuming pod identity, unsure.

CSI drivers are root on nodes.

Use service-api for now.

Want to give csi drivers their own credentials to do things on behalf of pods.


how to add identity to pod so when it requests cert through api add something to request so it can be determined
identity or request and if it has access to request cert.

#### Key Leanings / Take aways
Changes are going to resolve around the CSR request.


#### Action items
- Figure out what would be needed in certificates signing request  api to merge into sig-auth write down users stories 
and propose to sig-auth
