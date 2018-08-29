# K8s Node TLS bootstrapping via TPM

### Author: awly@

## Objective

This document proposes a mechanism to incorporate TPM evidence into the node
TLS bootstrapping process for use by a CSR approver in environments where TPMs
are available.

## Background

### Terminology

Trusted Platform Module (TPM) is a discrete hardware, firmware or software
device that (among other things) can store ECC and RSA keys and use them for
cryptographic operations.

Endorsement Key (EK) is a primary key for Endorsement Hierarchy baked into the
TPM.
EK Certificate is an x509 certificate created using EK and signed by an
external Certificate Authority (CA).

Attestation Identity Key (AIK) is a version of EK which is sign-only (as
opposed to decrypt-only for default EK).
AIK Certificate is EK Certificate with AIK as private key.

### Current TLS cert bootstrapping

[Node TLS
bootstrapping](https://kubernetes.io/docs/admin/kubelet-tls-bootstrapping/) is
a process of issuing client TLS certificates to a new k8s Node (kubelet) for
authorization to API server.

Current Kubelet TLS bootstrapping documentation suggests using a bootstrap
bearer token and either using the [“nodeclient” SubjectAccessReview
(SAR)](https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/certificates/approver/sarapprove.go)
approval flow or doing manual approvals. These present security issues:

* Weak bootstrap credentials: Bootstrap bearer tokens are bad for all the
  reasons bearer tokens are bad. They tend to be long lived (or static), and
  not single use. When used in combination with “nodeclient” SAR approval
  (discussed below), a compromised token allows an attacker join the cluster
  from arbitrary machines as a new Node.
* Weak approval flows:
  * “nodeclient” SAR approval is bad because it allows any node in the cluster
    to get credentials for any other node in the cluster.
  * Manual approval is bad because it doesn’t scale and operators make errors.

While other models of node TLS bootstrap are possible with the current
implementation, they are much more difficult to setup, undocumented and thus
not widely deployed.

### TPM

Private keys never leave the TPM boundaries in plain text.

TPMs can also generate keys and organize them in hierarchies, such that
parent-child relationship can be cryptographically proven.

Hardware TPMs are usually low-power chips with very limited storage and poor
performance. Software TPMs also have limited storage but can perform
cryptographic operations much faster.

## Overview

In the proposed design, the kubelet can use a TPM as proof of identity to API
server when creating TLS key pair used to authenticate to API server.

At machine creation time, the TPM can be provisioned with AIK Certificate
signed by a trusted CA. The certificate includes machine’s identifying info.
Kubelet’s TLS keys will be attested by AIK to link them to machine identity.

![Node bootstrap flow using TPM](node-tls-bootstrap-via-tpm.svg)

After machine setup, node bootstrapping will go as follows:

1. Kubelet generates an Elliptic Curve (EC) key pair and stores it on disk
1. Kubelet triggers AIK creation in the TPM
1. Kubelet requests that TPM certifies (signs) ECC public key with AIK
1. Kubelet loads TPM’s AIK cert into memory
1. Kubelet builds a Certificate Signing Request (CSR) using TLS key pair
1. Kubelet posts a CertificateSigningRequest object to API server including
   CSR, AIK cert and key certification
1. Master receives the CertificateSigningRequest
1. Master validates AIK cert, key certification and CSR contents
1. Master validates that client machine belongs to cluster (pluggable logic
   specific to cluster environment)
1. Master approves and signs the CSR
1. Kubelet observes signed CSR and stores the certificate (on disk or
   optionally in TPM’s NVRAM)

Some knowledge about expected cluster members on master is a prerequisite for
this design.

## Design

### TPM

#### AIK certificate

The device comes provisioned with an Attestation Identity Key (more
specifically a sign-only Endorsement Key) and x509 AIK Certificate signed by a
trusted CA.

OtherName field in SubjectAltName of AIK certificate contains identifying info
for the machine.

AIK is a sign-only restricted key - can only sign other keys loaded in the TPM.
We use this information (combined with proof of AIK ownership i.e. a signature)
as proof of machine’s identity to Master.

#### TPM ownership

TPM allows users to set authorization passwords on key hierarchies and
individual keys. Kubelet will initially be the exclusive user of the TPM and
will set a password on endorsement hierarchy. Password will be generated during
bootstrap and persisted in Kubelet-owned file on disk. This protects against
any future non-root user of the TPM in a container from being able to
impersonate the Kubelet.

#### TPM key attestation commands

1. TPM2_NV_Read - AIK template and nonce
1. TPM2_CreatePrimary - AIK, with loaded template
1. TPM2_LoadExternal TLS public key into NULL hierarchy
1. TPM2_Certify - AIK as signing key, ECC public key as subject
1. TPM2_FlushContext on TLS key to clean up
1. TPM2_NV_Read - load AIK cert from known index

### Kubelet

#### CSR creation

CSR is a PEM-encoded blob in CertificateSigningRequest API object. Kubelet will
need to append AIK certificate and proof of key ownership as extra PEM blocks
in the same field.

CSR creation steps:

1. create ECDSA key pair, store it on disk
1. generate the CSR in Kubelet’s memory
1. add an extension indicating extra TPM-based CSR data (OID, value “true”).
   This lets master distinguish between regular and TPM-attested CSRs.
1. sign ECC public key and metadata using AIK (described above)
1. build k8s CertificateSigningRequest object with Request field set to
   concatenated PEM blocks of: CSR, key certification, AIK certificate

### Example Master CSR validation

Steps to validate the CSR on master:

1. receive new CertificateSigningRequest object
1. parse PEM blocks from Request field (CSR, Public Key signature, AIK cert)
1. validate CSR contents
1. validate public key signature using AIK cert
1. validate AIK cert chain up to a known global CA (specified by master, not
   kubelet)
1. extract machine identity info from AIK cert
1. validate the machine identity against expected cluster members
   (provider/implementation-specific)
1. if all checks pass, approve and sign the CSR

The above validation can be implemented in certificate controller (a.k.a.
approver).

### Backwards-compatibility

The existing bootstrap token flow will remain supported. Approver will
differentiate between TPM-based and token-based CSRs via x509 extension (see
“CSR creation” section) to decide how to validate them.

Approver can choose to reject non-TPM CSRs.

### Future work

Below is a list of potential future improvements that build on the above
design.

#### Key pair stored in TPM

TPMs can generate and store key pairs and make them non-exportable: the private
key never leaves TPM API boundary.
[tls.Certificate.PrivateKey](https://golang.org/pkg/crypto/tls/#Certificate)
allows replacing in-memory key with a custom crypto.Signer implementation.

#### Certificate rotation

Certificate rotation currently uses existing certificate to request a new one.
If existing certificate expires before rotation succeeds (e.g. when node is
partitioned from cluster or suspended), the node requires operator
intervention.  Rotation can be changed to use the TPM-based flow, eliminating
the need for manual intervention.

#### Incorporate PCR quoting

PCR quotes can be added to the CSR. These quotes record the state of the
machine (firmware, kernel parameters, etc). Master may validate these values
during CSR approval.

## Implementation options

### Option 1: directly in Kubelet

#### New package `tpm`

Pckage `tpm` will have helper funcs for TPM-related functionality. Uses
[github.com/google/go-tpm/tpm2](https://github.com/google/go-tpm/tree/master/tpm2)
under the hood.

```go
package tpm

type TPM struct {...}
func Open(path) (*TPM, error)
func (t *TPM) Close() error

// AttestKey returns series of PEM blocks with attestation data.
func (t *TPM) AttestKey(privateKey) (attestData, error)
// VerifyCSR parses CSR PEM blocks with attestation data and verifies the
// signature (used by Approver).
func (t *TPM) VerifyCSR(csrData) error
```

#### New kubeconfig fields

New fields for TPM interaction are added under `users` in `kubeconfig`
([AuthInfo
struct](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/client-go/tools/clientcmd/api/types.go)):

1. BootstrapCSRAttestation={none|tpm} - extra attestation mode for CSRs. none
   is default. Field is behind a [feature
   gate](https://github.com/kubernetes/community/blob/master/contributors/devel/api_changes.md#alpha-field-in-existing-api-version).
   Note: other values may be added in the future for alternative attestation
   schemes.
1. TPM (new type TPMConfig) - TPM-specific config fields
1. TPM.Path string - path to TPM device (character device or unix socket)
1. TPM.AIKCertIndex int - NVRAM index of AIK certificate
1. TPM.AIKTemplateIndex int - NVRAM index of AIK template

#### Kubelet code

[bootstrap.LoadClientCert](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/certificate/bootstrap/bootstrap.go#L47)
is the entry point.

Split [csr.RequestNodeCertificate](https://github.com/kubernetes/client-go/blob/master/util/certificate/csr/csr.go#L51) into

```go
func MakeNodeCSR(privateKey, nodeName) (csrData, error)
func RequestNodeCertificate(client, csrData, privateKey) (certData, error)
```

In `bootstap.LoadClientCert`:

```
csrData, err := csr.MakeNodeCSR(privateKey, nodeName)
// handle err
switch config.BootstrapCSRAttestation {
    case "tpm":
        attestData, err := tpm.AttestKey(privateKey)
        // handle err
        csrData = append(csrData, attestData...)
}
certData, err := csr.RequestNodeCertificate(client, csrData, privateKey)
// handle err
```

Note: in this doc `privateKey` is always `crypto.PrivateKey`, not a PEM-encoded
byte blob.

### Option 2: exec authentication plugin

See [design for exec auth
plugin](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/auth/kubectl-exec-plugins.md)

The plugin hides all the TPM interaction and produces private key and signed
certificate.

The certificate is cached on disk. When existing certificate is about to expire
(e.g. <1d), exec plugin requests a new one and overwrites it.
