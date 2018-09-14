# Providing external secrets to Pods in Kubernetes

Status: Pending

Version: Alpha

Implementation Owner: sig-auth

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
secret at the time it is actually accessed by the container. On the
(remote) secrets provider side (vault), access control decision,
logging, or on-the-fly generation of secrets can happen, while taking
into account the identity of the Node and Pod that request the
particular secret.

## Proposal

We are proposing flexVolume plugin which would utilize (Filesystem in
Userspace) FUSE. It allows creating virtual filesystem where the
content can be produced on-the-fly and behaviour of the individual
entries can be actively controlled. The individual secrets would
appear as files in the filesystem, while being generated (or even
renewed, in case of certificates) via the backend logic, if needed.

The FUSE filesystem is represented by a mountpoint of type `fuse.*`
and backed by a process, run on the Node with `root` privileges.
Therefore, it can access per-Node credentials like Kerberos keytabs,
tokens, or TLS client certificates, for authenticating to remote,
external vault / external secrets provider.

During flexVolume `mount` operation, parameters like (formatted
for clarity)

```
/var/lib/kubernetes/volumes/pods/215904af-a29b-11e7-a06b-5254005fe346/volumes/example.com~secrets-fuse/secrets-fuse
{
    "kubernetes.io/fsGroup":"1000020000",
    "kubernetes.io/fsType":"",
    "kubernetes.io/pod.name":"test-pod",
    "kubernetes.io/pod.namespace":"default",
    "kubernetes.io/pod.uid":"215904af-a29b-11e7-a06b-5254005fe346",
    "kubernetes.io/pvOrVolumeName":"hashicorp-cli",
    "kubernetes.io/readwrite":"rw",
    "kubernetes.io/serviceAccount.name":"default"
}
```

are passed to the flexVolume plugin. The Pod identity information,
namely `kubernetes.io/pod.uid`, `kubernetes.io/pod.namespace`,
and `kubernetes.io/pod.name` are present and can be passed to the
external vault to request secrets for a particular Pod.

The network protocols, data formats, and authentication mechanisms
differ between vault / secrets provider. Implementing them in the
FUSE flexVolume plugin would require different plugins for different
vaults. Therefore, the actual operation of secret retrieval is
delegated to an external "helper" program, which in two dozens lines
of shell code can capture the specifics of the vault and the use
vault-specific command line or API tool. The helper program can also
be used to configure the FUSE, namely specify which paths should be
presented and accessible as directories, and what `mount` parameters
should be passed as parameters during the secret entry retrieval.

The flexVolume plugins need to be placed into
/usr/libexec/kubernetes/kubelet-plugins/volume/exec/ subdirectory,
to be discovered by Kubernetes. We propose for the helper program
to be placed into
/usr/libexec/kubernetes/kubelet-plugins/volume/libexec/, in the
subdirectory and with the basename matching the flexVolume plugin.

For example, if the flexVolume plugin is named
`example.com/hashicorp-cli-fuse`, the plugin itself (or a symlink to
its binary) has to be in
`/usr/libexec/kubernetes/kubelet-plugins/volume/exec/example.com~hashicorp-cli-fuse/hashicorp-cli-fuse`
and the matching helper will be looked up in
`/usr/libexec/kubernetes/kubelet-plugins/volume/libexec/example.com~hashicorp-cli-fuse/hashicorp-cli-fuse`.
For different vault implementation, for example
`example.com/custodia-cli-fuse`, the same flexVolume plugin binary
can be used and symlinked from
`/usr/libexec/kubernetes/kubelet-plugins/volume/exec/example.com~custodia-cli-fuse/custodia-cli-fuse`,
with the helper in
`/usr/libexec/kubernetes/kubelet-plugins/volume/libexec/example.com~custodia-cli-fuse/custodia-cli-fuse`.

During the `mount` operation of the flexVolume plugin, the helper
external program is invoked with `mount` parameter and it should
produce JSON object/hash on its standard output with configuration.

Key `enable-dirs` can be an array of strings of directories that
should be "present" on the FUSE filesystem. It defaults to `"/"`.

Key `mount-param` can be an array of strings of parameters
from the flexVolume plutin `mount` second parameter, the JSON flags.
It defaults to empty list.

For example, bash helper snippet

```
case "$1" in
        mount)
                echo "{"
                echo "\"enable-dirs\": [ \"/db\" ],"
                echo "\"mount-param\": [ \"kubernetes.io/pod.namespace\", \"kubernetes.io/pod.name\" ],"
                echo "}"
                ;;
```

will cause the access to "directory" `db/` under the mountpoint
of the flexVolume to be allowed and the values of Pod namespace
and Pod name to be passed as parameters besides the path of the
secret on the filesystem during the actual retrieval of the secret
entry.

If the flexVolume filesystem in mounted in the container under
`/mnt/secrets-external`, opening and reading
`/mnt/secrets-external/db/password` will cause the helper to be
invoked with parameters `get`, `/db/password`, plus the Pod namespace
and Pod name, per the `mount-param` configuration returned during
the `mount` operation. The standard output of the `get` operation
is then presented as `/mnt/secrets-external/db/password` file. If the
retrieval fails, for example because the secret is not present in
the external vault or there was an error, the file entry will be
missing.

The last accessed secret is cached by the FUSE driver to avoid
refetching the value from the external vault between getattr, open,
and read operations.

## User Experience

### Use Cases

* Retrieval of database password from HashiCorp Vault.

* Access to TLS certificate for the Web application running in the
  Pod / container; the certificate can be either in FreeIPA Vault
  or generated or renewed on-the-fly via FreeIPA CA; Custodia is used
  as the API.

Let's assume the binary of the flexVolume plugin is placed to
`/usr/bin/flexvolume-fuse-external`. To enable it as
`example.com/hashicorp-cli-fuse`, make symlink to it from
`/usr/libexec/kubernetes/kubelet-plugins/volume/exec/example.com~hashicorp-cli-fuse/hashicorp-cli-fuse`,
and on the Nodes also create (or symlink) vault-specific helper
`/usr/libexec/kubernetes/kubelet-plugins/volume/libexec/example.com~hashicorp-cli-fuse/hashicorp-cli-fuse`.
Example [hashicorp-cli-helper](hashicorp-cli-helper.sh) which invokes
`hashicorp-vault read` is included.

The same flexVolume plugin binary can be symlinked for example from
`/usr/libexec/kubernetes/kubelet-plugins/volume/exec/example.com~custodia-cli-fuse/custodia-cli-fuse`
to enable it as flexVolume plugin `example.com/custodia-cli-fuse`. The
helper is then expected in
`/usr/libexec/kubernetes/kubelet-plugins/volume/libexec/example.com~custodia-cli-fuse/custodia-cli-fuse`.
See [custodia-cli-helper](custodia-cli-helper.sh) for an example of
helper calling `custodia-cli read`.

In these examples, `$3` and `$4` command line parameters will be
the values of `kubernetes.io/pod.namespace` and `kubernetes.io/pod.name`
and the helpers will use them to construct the path of the secret
to be retrieved. However, it could also pass them as HTTP headers
to `curl` or do any other manipulation, to match the external vault
or secrets provider setup.

Of coure, the helper could be written in any language and could
make the network operations against the external vault directly, as
long as it produces the JSON configuration when invoked with `mount`
parameter, and the secret content when invoked with `get` operation.

The FUSE filesystem can get mounted to Pod with

```
apiVersion: v1
kind: Pod
metadata:
  name: test-pod-hashicorp
spec:
  containers:
  - image: registry.access.redhat.com/rhel7
    name: test-pod-hashicorp
    command: ["/usr/bin/sleep"]
    args: ["infinity"]
    volumeMounts:
    - mountPath: /mnt/secrets-external
      name: hashicorp-cli
  volumes:
  - flexVolume:
      driver: example.com/hashicorp-cli-fuse
    name: hashicorp-cli
```

Within the container of the pod, the application can then access
`/mnt/secrets-external/db/password`, the FUSE filesystem will see
it as access to `/db/password` (provided the helper enabled `/db`
in its `enabled-dirs` list) and it will invoke the helper to get the
secret and present it as the file `/mnt/secrets-external/db/password`.

Since the FUSE provider runs as root, it will have access to necessary
credentials to prove its own identity to the remote provider. In the
[custodia-cli-helper](custodia-cli-helper.sh) example, the
`CUSTODIA_CLIENT_KEYTAB` value could be set to `/etc/krb5.keytab`
in `/etc/custodia/custodia-cli.conf` which is sourced by the helper,
to use the identity of the `host/` for GSSAPI operation. In the
[hashicorp-cli-helper](hashicorp-cli-helper.sh) example, we assume
that `/etc/hashicorp/hashicorp-cli.conf` contains `VAULT_TOKEN` value
or `VAULT_CLIENT_CERT` path.

In production setup, we expect that the remote vault or secrets
provider will make verification against Kubernetes API, checking for
example that the Pod for which the Node is requesting the secret,
is indeed currently scheduled and running on that Node. That way, if
a Node is compromised, only (secrets of) Pods running on that single
Node will be compromised.

## Implementation

The [flexvolume-fuse-external.c](flexvolume-fuse-external.c) source
code is FUSE `flexVolume` plugin which uses `libfuse` and `json-c`
libraries to implement userspace filesystem.

The compilation of the plugin/FUSE filesystem is specific to operating
system on which the Kubernetes cluster is deployed. For example, on
RHEL, CentOS, and Fedora, `gcc`, `fuse-devel`, and `json-c-devel`
packages need to be installed (presumably via `yum` or `dnf`). Then
the source code can be compiled with

```
gcc -Wall flexvolume-fuse-external.c $( pkg-config fuse json-c --cflags --libs ) \
      -o flexvolume-fuse-external
```

During compilation, `-D LOG_FILE=<path-to-log-file>` can be used to
enable logging of the plugin invocations and parameters passed to it.
Compile-time option `-D DEBUG=1` enables verbose logging to the log file.

Note that The flexVolume's `init` has to return `"selinuxRelabel": false` to
prevent labelling operation which is not supported on the FUSE
filesystem.

The FUSE filesystem has to be mounted with `allow_other` option to
allow non-root uids (that the containers run as) to access the filesystem.

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
* https://www.vaultproject.io/docs/commands/read-write.html
* https://github.com/latchset/custodia

