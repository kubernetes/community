# Kubelet Server Certificate Bootstrap & Rotation

## Overview

Currently, when a kubelet first starts, it generates a self signed
certificate/key pair that is used for accepting incoming TLS connections. This
proposal covers a process for generating a key locally and then issuing a
Certificate Signing Request to the cluster API server to get an associated
certificate signed by the cluster Certificate Authority. Also, as certificates
approach expiration, the same mechanism will be used to request an updated
certificate.

## Use Cases

1. Kubelet starts for the first time on a new node.
2. Kubelet certificate approaches expiration.

## Proposed Design

### During Kubelet Boot Sequence

1. Look on disk for an existing cert/key pair.
1. If there is an existing cert/key pair, load them.
1. If there is no existing cert/key pair, at the point where the kubelet
   currently generates a key and certificate, the kubelet will:
    - generate a key
    - create a certificate signing request
    - request a signed certificate from the API server
    - wait for a response
    - Store the new cert/key pairs in the directory specified by `--cert-dir`,
      or the default value if it is unspecified.

### As Expiration Approaches

Rotating a Kubelet server certificate will work by generating a new private key,
issuing a new Certificate Signing Request to the API Server, safely updating the
cert/key pair on disk, begin using the new cert/key pair.
 
1. Store the new cert/key pairs in the directory specified by `--cert-dir`, or the
   default value if it is unspecified. This will allow the kubelet to have a
   place for storing the multiple cert/key pairs that it might have available at
   any given moment (because of a rotation in progress).
    - When cert/key data is directly embedded in the kubelet config file
      (client-certificate-data, client-key-data)
        - The values be used only if there are no updated cert/key pairs due to
          rotation
    - When cert/key files are specified in kubeconfig, or with the
      `--tls-cert-file` and `--tls-private-key-file`:
        - These values be used only if there are no updated cert/key pairs due
          to rotation
1. There will be a feature gate to enable certificate bootstrapping and
   rotation, `RotateKubeletServerCertificate`
1. Centralize certificate access within the kubelet code to the
   CertificateManager. The CertificateManager will be responsible for:
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
          kubelet-server-<timestamp>.pem.
        - Write kubelet-server-updated.pem symlink to point to the new cert/key
          pair.
        - Delete kubelet-server-current.pem
        - Move kubelet-server-updated.pem to kubelet-server-current.pem
1. Currently the Certificate Request Signing API is hard coded to issue
   certificates for 1 year. Making this a parameter that can be set quite short
   during end to end testing would enable easy certificate rotation testing.

### Certificate Approval

With the kubelet requesting certificates be signed as part of is boot sequences,
and on an ongoing basis, it makes cluster administration easier if certificate
signing requests from the kubelet were automatically approved. A full solution
to this will require authenticating the kubelet making the request in some way,
and will be covered by a [future design]
(https://github.com/kubernetes/kubernetes/issues/45030).

