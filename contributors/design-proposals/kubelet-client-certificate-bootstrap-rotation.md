# Kubelet Client Certificate Bootstrap & Rotation

## Overview

Currently, a kubelet has a certificate/key pair to use for making connection to
the cluster master that authenticates the kubelet to the master. The certificate
is supplied to the kubelet when it is first booted, via an out of cluster
mechanism. This proposal covers a process for rotating the initial cert/key pair
as expiration of the certificate approaches.  The existing cert/key pair are
used to authenticate to the cluster API server and issue a request to the
Certificate Signing Request API to sign a new key.

## Use Cases

1. Kubelet starts for the first time on a new node.
2. Kubelet certificate approaches expiration.

## Proposed Design

### During Kubelet Boot Sequence

1. Look on disk for an existing cert/key pair managed by the certificate
   manager.
1. If there is an existing cert/key pair, load them.
1. If there is no existing cert/key pair, look for cert/key data specified in
   the kubelet config file:
    - values encoded in the kubeconfig (CertData, KeyData, CAData)
    - file references in the kubeconfig (CertFile, KeyFile, CAFile)
1. If the certificate is a bootstrap certificate, use the certificate to:
    - generate a key
    - create a certificate signing request
    - request a signed certificate from the API server
    - wait for a response
    - Store the new cert/key pairs in the directory specified by `--cert-dir`,
      or the default value if it is unspecified.

### As Expiration Approaches

Rotating a Kubelet client certificate will work by generating a new private key,
issuing a new Certificate Signing Request to the API Server, safely updating the
cert/key pair on disk, begin using the new cert/key pair.

1. Store the new cert/key pairs in the directory specified by `--cert-dir`, or the
   default value if it is unspecified. This will allow the kubelet to have a
   place for storing the multiple cert/key pairs that it might have available at
   any given moment (because of a rotation in progress).
    - When cert/key files are specified in the kubeconfig, these will be used if
      a newer, rotated, cert/key pair does not exist.
1. There will be a feature gate to enable certificate bootstrapping and
   rotation, `RotateKubeletClientCertificate`
1. Centralize certificate access within the kubelet code to the
   CertificateManager. Whenever rotation is enabled, CertificateManager will be
   responsible for:
    - Providing the correct certificate to use for establishing TLS connections.
    - Generating new private keys and requesting new certificates when the
      current certificate approaches expiry.
    - Since certificates can rotate at any time, all other parts of the kubelet
      should ask the CertificateManager for the correct certificate each time a
      certificate is used. No certificate caching except by the
      CertificateManager.
        - TLS connections should prefer to set the
          [`GetCertificate`](https://golang.org/pkg/crypto/tls/#Config) and
          [`GetClientCertificate`](https://golang.org/pkg/crypto/tls/#Config)
          callbacks so that the connection dynamically requests the certificate
          as needed.
    - Recovering from kubelet crashes or restarts that occur while certificate
      transitions are in flight (request issued, but not yet signed, etc)
1. The RBAC role for nodes will be updated to allow nodes to request new
   certificates.
1. Have the CertificateManager repeat the CSR process as certificate expiration
   approaches.
    - New certificates will be requested when the configured duration threshold
      has been exceeded.
    - Crash safe file structure
        - Private key for requests in flight are held in memory. In case of
          interruption, they will be abandoned.
        - When the corresponding signed certificate is received, the cert/key
          pair will be written to a single file, eg.
          kubelet-client-<timestamp>.pem.
        - Write kubelet-client-current.pem symlink to point to the new cert/key
          pair.

### Certificate Approval

With the kubelet requesting certificates be signed as part of is boot sequences,
and on an ongoing basis, certificate signing requests from the kubelet need to
be auto approved to make cluster administration manageable. Certificate signing
request approval is complete, and covered by [this design]
(https://github.com/kubernetes/kubernetes/issues/45030).

