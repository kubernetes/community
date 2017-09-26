# Providing external secrets to Pods in Kubernetes

Status: Pending

Version: Alpha

Implementation Owner: openshift/sig-security

## Motivation

Kubernetes support the Secrets mechanism where secrets are stored in
etcd and made available in Pods and containers via `volumeMounts` or
environment variables. In these cases, the secrets are either
automatically created in etcd (credentials with access to API) or
generated and stored by admins/users.

In some cases though, secrets can primarily be created and stored
in sources external to the Kubernetes cluster. For example, TLS
certificates can be created, signed, and renewed by external
certificate authority like FreeIPA/IdM's CA, Let's Encrypt, or AD CS,
and they can be stored in external vault storage like FreeIPA/IdM Vault
or HashiCorp Vault. To make them available via Kubernetes's Secrets
mechanism, we'd need to have the certificates created beforehand and
stored in etcd first. That might not be the ideal workflow because it
does not support requirements like automatically renewing the
certificate if the copy found in vault is getting close to expiration,
or auditing the actual access of individual containerized applications
to the secrets.

It is possible to run per-Node daemon via DaemonSet with `hostPath`
access to per-Node credentials (to authenticate to the external
vault / secrets provider), exposing Unix socket on the filesystem of
the Node. Using flexVolume mechanism, the directory with that Unix
socket can then be made available to Pods and containers that need
access to the secrets. The daemon in the DaemonSet can then use
`SO_PEERCRED`/`SO_PEERSEC` to get the identity of the process accessing
the service via the Unix socket, and authorize it based on the Pod
identity and other attributes / annotations which can be retrieved
via Kubernetes API.

One drawback of this approach is the need for active client tool or
library in each of the container, to invoke the daemon's API over
the Unix socket to request the appropriate secret / certificate.

For that reason we are looking for mechanism of exposing the secrets
from external sources to Pods and their containers via filesystem
interface. With that approach, the individual secrets would be seen as
files in filesystem structure, while backed by active mechanisms
of external secrets providers or vaults, and would only request the
secret at the time it is actually accessed by the container.

## Proposal

We are proposing flexVolume plugin which would utilize FUSE. FUSE
allows creating virtual filesystem where the content can be produced
on-the-fly and behaviour of the individual entries can be actively
controlled. The individual secrets would appear as files in the
filesystem, while being generated (or even renewed, in case of
certificates) via the backend logic, if needed.

During `mount` operation, parameters like (formatted for clarity)

```
/var/lib/kubernetes/volumes/pods/215904af-a29b-11e7-a06b-5254005fe346/volumes/example.com~secrets-fuse/secrets-fuse
{
    "kubernetes.io/fsGroup":"1000020000",
    "kubernetes.io/fsType":"",
    "kubernetes.io/pod.name":"test-pod",
    "kubernetes.io/pod.namespace":"default",
    "kubernetes.io/pod.uid":"215904af-a29b-11e7-a06b-5254005fe346",
    "kubernetes.io/pvOrVolumeName":"secrets-external",
    "kubernetes.io/readwrite":"rw",
    "kubernetes.io/serviceAccount.name":"default"
}
```

are passed to the flexVolume plugin, so the flexVolume plugin can
parse the parameters and get the `kubernetes.io/pod.uid` value
and pass it to the FUSE filesystem via `mount` option. That provides
the FUSE filesystem with the identity of the clients that will be
accessing the secrets which can be used for access authorization.

The FUSE filesystem runs as root so it can have access to all needed
credentials (certificates, Kerberos keytabs) needed to further
authenticate to the external vault / external secrets provider.

## User Experience

### Use Cases

* Access to TLS certificate for the Web application running in the
  Pod / container; the certificate can be either in FreeIPA Vault
  or generated or renewed on-the-fly via FreeIPA CA; Custodia is used
  as the API.

Let's assume a FUSE filesystem which provides access to its individual
files as requests for certificates is implemented as `/usr/bin/secrets-fuse`.

With (trivial for clarity) flexVolume plugin

```
#!/bin/bash

# Put to
# /usr/libexec/kubernetes/kubelet-plugins/volume/exec/example.com~secrets-fuse/secrets-fuse

case "$1" in
	init)
		echo '{ "status": "Success", "capabilities": {"attach": false, "selinuxRelabel": false}}'
	;;
	mount)
		mkdir -p "$2"
		# TODO: parse kubernetes.io/pod.uid and pass it
		mount -t fuse -o allow_other /usr/bin/secrets-fuse "$2"
		echo '{ "status": "Success" }'
	;;
	unmount)
		umount -f "$2"
		rmdir "$2"
		echo '{ "status": "Success" }'
	;;
	*)
		echo '{ "status": "Not supported" }'
		exit 1
esac
```

the FUSE filesystem can get mounted to Pod like

```
apiVersion: v1
kind: Pod
metadata:
  name: test-pod
spec:
  containers:
  - image: registry.access.redhat.com/rhel7
    imagePullPolicy: Never
    name: test-pod
    volumeMounts:
    - mountPath: /mnt/secrets-external
      name: secrets-external
  volumes:
  - flexVolume:
      driver: example.com/secrets-fuse
    name: secrets-external
```

Within the container of the pod, the application can then access
`/mnt/secrets-external/HTTP/www.example.com` and the FUSE filesystem
will see this as access to `HTTP/www.example.com` file. Via the
`kubernetes.io/pod.uid`, it can verify (against its internal rules
or for example via Kubernetes API against attributes of that Pod)
whether the Pod should be allowed access to the `HTTP/www.example.com`
secret, and then invoke tools like `custodia-cli` via

```
custodia-cli get HTTP/www.example.com
```

to get the secret from remote, external source. Alternatively, the
Pod identity with other attributes / annotations can be passed to the
remote provider to run the authorization on the remote side.

Since the FUSE provider runs as root, it will have access to necessary
credentials to prove its own identity to the remote provider.

In production setup, we expect that remote provider to be a Custodia
server which will make additional verification against Kubernetes API,
checking that the on-Node identity which makes the request for
secrets is only making the request on behalf of the Pods that are
currently running on the Node. That way, if a Node is compromised,
only (secrets of) Pods running on that single Node are compromised.
The Custodia server on that separate host will then proxy the requests
to the actual vault or secrets provider.

## Implementation

Implementation of FUSE filesystem using libfuse 2.9 will show example
invocation of `custodia-cli` to get secrets from remote Custodia server.
Accompanying flexVolume driver would then mount the FUSE filesystem
when configured for Pod.

The flexVolume's `init` has to return `"selinuxRelabel": false` to
prevent labelling operation which is not supported on the FUSE
filesystem.

The FUSE filesystem has to be mounted with `allow_other` option to
allow non-root uids (that the containers run as) to access the filesystem.

The `kubernetes.io/pod.uid` of `mount` parameters needs to be parsed.

During `getattr` and `read`, `struct fuse_context *fuse_get_context(void)`
can be called to get pid, uid, and git of the calling process. This
information can be used to cross-verify it against the the Pod
identity to prevent leaking the information from the mounted
filesystem on the host, for example during backup.

The SELinux boolean `virt_sandbox_use_fusefs` has to be set to true to
allow Kubernetes containers to access the filesystem.

### Client/Server Backwards/Forwards compatibility

Not affected.

## Alternatives considered

Container Storage Interface approach might be the long-run preferred
method but at this point flexVolume implementation will allow immediate
use.

## References

* https://github.com/kubernetes/community/blob/master/contributors/devel/flexvolume.md
* https://github.com/libfuse/libfuse/tree/fuse-2_9_bugfix
* https://github.com/latchset/custodia
* https://github.com/latchset/custodia.openshift

