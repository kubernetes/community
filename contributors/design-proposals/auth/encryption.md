# Encryption

## Abstract

The scope of this proposal is to ensure that resources can be encrypted at the
datastore layer with sufficient metadata support to enable integration with
multiple encryption providers and key rotation. Encryption will be optional for
any resource, but will be used by default for the Secret resource. Secrets are
already protected in transit via TLS.

Full disk encryption of the volumes storing etcd data is already expected as
standard security hygiene.  Adding the proposed encryption at the datastore
layer defends against malicious parties gaining access to:

- etcd backups; or
- A running etcd instance without access to memory of the etcd process.

Allowing sensitive data to be encrypted adheres to best practices as well as
other requirements such as HIPAA.

## High level design

Before a resource is written to etcd and after it is read, an encryption
provider will take the plaintext data and encrypt/decrypt it.  These providers
will be able to be created and turned on depending on the users needs or
requirements and will adhere to an encryption interface.  This interface will
provide the abstraction to allow various encryption mechanisms to be
implemented, as well as for the method of encryption to be rotated over time.

For the first iteration, a default provider that handles encryption in-process
using a locally stored key on disk will be developed.

## Kubernetes Storage Changes

Kubernetes requires that an update that does not change the serialized form of
object not be persisted to etcd to prevent other components from seeing no-op
updates.

This must be done within the Kubernetes storage interfaces - we will introduce a
new API to the Kube storage layer that transforms the serialized object into the
desired at-rest form and provides hints as to whether no-op updates should still
persist (when key rotation is in effect).

```go
// ValueTransformer allows a string value to be transformed before being read from or written to the underlying store. The methods
// must be able to undo the transformation caused by the other.
type ValueTransformer interface {
        // TransformFromStorage may transform the provided data from its underlying
        // storage representation or return an error.  Stale is true if the object
        // on disk is stale (encrypted with an older key) and a write to etcd
        // should be issued, even if the contents of the object have not changed.
        TransformFromStorage([]byte) (data []byte, stale bool, err error)
        // TransformToStorage may transform the provided data into the appropriate form in storage or return an error.
        TransformToStorage([]byte) (data []byte, err error)
}
```

When the storage layer of Kubernetes is initialized for some resource, an
implementation of this interface that manages encryption will be passed down.
Other resources can use a no-op provider by default.

## Encryption Provider

An encryption provider implements the ValueTransformer interface. Out of the box
this proposal will implement encryption using a standard AES-GCM performing
AEAD, using the standard Go library for AES-GCM.

Each encryption provider will have a unique string identifier to ensure
versioning of the ciphertext in etcd, and to allow future schemes to be added.

During encryption, only a single provider is required.  During decryption,
multiple providers or keys may be in use (when migrating from an older version
of a provider, or when rotating keys), and thus the ValueTransformer
implementation must be able to delegate to the appropriate provider.

Note that the ValueTransformer is a general storage interface and not related to
encryption directly. The AES implementation linked below combines
ValueTransformer and encryption provider.

### AES-GCM Encryption provider

Implemented in [#41939](https://github.com/kubernetes/kubernetes/pull/41939).

The simplest possible provider is an AES-GCM encrypter/decrypter using AEAD,
where we create a unique nonce on each new write to etcd, use that as the IV for
AES-GCM of the value (the JSON or protobuf data) along with a set of
authenticated data to create the ciphertext, and then on decryption use the
nonce and the authenticated data to decode.

The provider will be assigned a versioned identifier to uniquely pair the
implementation with the data at rest, such as “k8s-aes-gcm-v1”.  Any
implementation that attempts to decode data associated with this provider id
must follow a known structure and apply a specific algorithm.

Various options for key generation and management are covered in the following
sections. The provider implements one of those schemes to retrieve a set of
keys. One is identified as the write key, all others are used to decrypt data
from previous keys.  Keys must be rotated more often than every 2^32 writes.

The provider will use the recommended Go defaults for all crypto settings
unless otherwise noted.  We should use AES-256 keys (32 bytes).

Process for transforming a value (object encoded as JSON or protobuf) to and
from stable storage will look like the following:

Layout as written to etcd2 (json safe string only):
```
NONCE := read(/dev/urandom)
PLAIN_TEXT := <VALUE>
AUTHENTICATED_DATA := ETCD_KEY
CIPHER_TEXT := aes_gcm_encrypt(KEY, IV:NONCE, PLAIN_TEXT, A:AUTHENTICATED_DATA)
BASE64_DATA := base64(<NONCE><CIPHER_TEXT>)
STORED_DATA := <PROVIDER>:<KEY_ID>:<BASE64_DATA>
```

Layout as written to etcd3 (bytes):
```
NONCE := read(/dev/urandom)
PLAIN_TEXT := <VALUE>
AUTHENTICATED_DATA := ETCD_KEY
CIPHER_TEXT := aes_gcm_encrypt(KEY, IV:NONCE, PLAIN_TEXT, A:AUTHENTICATED_DATA)
STORED_DATA := <PROVIDER_ID>:<KEY_ID>:<NONCE><AUTHENTICATED_DATA><CIPHER_TEXT>
```

Pseudo-code for encrypt (golang):
```go
block := aes.NewCipher(primaryKeyString)
aead := cipher.NewGCM(c.block)
keyId := primaryKeyId

// string prefix chosen over a struct to minimize complexity and for write
// serialization performance.
// for each write
nonce := make([]byte, block.BlockSize())
io.ReadFull(crypto_rand.Reader, nonce)
authenticatedData := ETCD_KEY
cipherText := aead.Seal(nil, nonce, value, authenticatedData)
storedData := providerId + keyId + base64.Encode(nonce + authenticatedData + cipherText)
```

Pseudo-code for decrypt (golang):
```go
// for each read
providerId, keyId, base64Encoded := // slice provider and key from value

// ensure this provider is the one handling providerId
aead := // lookup an aead instance for keyId or error
bytes := base64.Decode(base64Encoded)
nonce, authenticatedData, cipherText := // slice from bytes
out, err := aead.Open(nil, nonce, cipherText, authenticatedData)
```

### Alternative Considered: SecretBox

Using [secretbox](https://godoc.org/golang.org/x/crypto/nacl/secretbox) would
also be a good choice for crypto. We decided to go with AES-GCM for the first
implmentation since:

- No new library required.
- We'd need to manage AEAD ourselves.
- The cache attack is not much of a concern on x86 with AES-NI, but is more so
  on ARM

There's no problem with adding this as an alternative later.

## Configuration

We will add the following options to the API server. At API server startup the
user will specify:

```yaml
--encryption-provider-config=/path/to/config
--encryption-provider=default
--encrypt-resource=v1/Secrets
```

The encryption provider will check it has the keys it needs and if not, generate
them as described in the following section.

## Key Generation, Distribution and Rotation

To start with we want to support a simple user-driven key generation,
distribution and rotation scheme. Automatic rotation may be achievable in the
future.

To enable key rotation a common pattern is to have keys used for resource
encryption encrypted by another set of keys (Key Encryption Keys aka KEK).  The
keys used for encrypting kubernetes resources (Data Encryption Keys, aka DEK)
are generated by the apiserver and stored encrypted with one of the KEKs.

In future versions, storing a KEK off-host and off-loading encryption/decryption
of the DEK to AWS KMS, Google Cloud KMS, Hashicorp Vault etc. should be
possible. The decrypted DEK would be cached locally after boot.

Using a remote encrypt/decrypt API offered by an external store will be limited
to encrypt/decrypt of keys, not the actual resources for performance reasons.

Incremental deliverable options are presented below.

### Option 1: Simple list of keys on disk

In this solution there is no KEK/DEK scheme, just single keys in a list on disk.
They will live in a file specified by the --encryption-provider-config, which
can be an empty file when encryption is turned on.

If the key file is empty or the user calls PUT on a /rotate API endpoint keys
are generated as follows:

1. A new encryption key is created.
1. The key is added to a file on the API master with metadata including an ID
   and an expiry time. Subsequent calls to rotate will prepend new keys to the
   file such that the first key is always the key to use for encryption.
1. The list of keys being used by the master is updated in memory so that the
   new key is in the list of read keys.
1. The list of keys being used by the master is updated in memory so that the
   new key is the current write key.
1. All secrets are re-encrypted with the new key.

Pros:

 - Simplicity.
 - The generate/write/read interfaces can be pluggable for later replacement
   with external secret management systems.
 - A single master shouldn't require API Server downtime for rotation.
 - No unseal step on startup since the file is already present.
 - Attacker with access to /rotate is a DoS at worst, it doesn't return any
   keys.

Cons:

 - Coordination of keys between a deployment with multiple masters will require
   updating the KeyEncryptionKeyDatabase file on disk and forcing a re-read.
 - Users will be responsible for backing up the keyfile from the API server
   disk.

### Option 2: User supplied encryption key

In this solution there is no KEK/DEK scheme, just single keys managed by the
user. To enable encryption a user specifies the "user-supplied-key" encryption
provider at api startup. Nothing is actually encrypted until the user calls PUT
on a /rotate API endpoint:

1. A new encryption key is created.
1. The key is provided back to the caller for persistent storage. Within the
   cluster, it only lives in memory on the master.
1. The list of keys being used by the master is updated in memory so that the
   new key is in the list of read keys.
1. The list of keys being used by the master is updated in memory so that the
   new key is is the current write key.
1. All secrets are re-encrypted with the new key.

On master restart the api server will wait until the user supplies the list of
keys needed to decrypt all secrets in the database. In most cases this will be a
single key unless the re-encryption step was incomplete.

Pros:

 - Simplicity.
 - A single master shouldn't require API Server downtime.
 - User is explicitly in control of managing and backing up the encryption keys.

Cons:

 - Coordination of keys between a deployment with multiple masters is not
   possible. This would have to be added as a subsequent feature using a
   consensus protocol.
 - API master needs to refuse to start and wait on a decrypt key from the user.
 - /rotate API needs to be strongly protected: if an attacker can cause a
   rotation and get the new key, it might as well not be encrypted at all.

### Option 3: Encrypted DEKs in etcd, KEKs on disk

In order to take an API driven approach for key rotation, new API objects (not
exposed over REST) will be defined:

* Key Encryption Key (KEK) - key used to unlock the Data Encryption Key. Stored
  on API server nodes.
* Data Encryption Key (DEK) - long-lived secret encrypted with a KEK. Stored in
  etcd encrypted. Unencrypted in-memory in API servers.
* KEK Slot - to support KEK rotation there will be an ordered list of KEKs
  stored in the KEK DB. The current active KEK slot number, is stored in etcd
  for consistency.
* KEK DB - a file with N KEKs in a JSON list. KEK[0], by definition, is null.

```go
type DataEncryptionKey struct {
    ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
    Value string    // Encrypted
}
```

```go
type KeyEncryptionKeySlot struct {
    ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
    Slot int
}
```

```go
type KeyEncryptionKeyDatabase struct {
    metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
    Keys []string
}
```

To enable encryption a user must first create a KEK DB file and tell the API
server to use it with `--encryption-provider-config=/path/to/config`.  The
file will be a simple YAML file that lists all of the keys:

```yaml
kind: KeyEncryptionKeyDatabase
version: v1
keys:
 - foo
 - bar
 - baz
```

The user will also need to specify the encryption provider and the resources to
encrypt as follows:
```yaml
--encryption-provider-config=/path/to/key-encryption-key/db
--encryption-provider=default
--encrypt-resource=v1/Secrets
--encrypt-resource=v1/ConfigMap
```

Then a user calls PUT on a /rotate API endpoint the first time:

1. A new encryption key (unencrypted DEK) is created.
1. Encrypt DEK with KEK[1]
1. The list of DEKs being used by the master is updated in memory so that the
   new key is in the list of read keys.
1. The list of DEKs being used by the master is updated in etcd so that the
   new key is in the list of read keys available to all masters.
1. Confirm that all masters have the new DEK for reading. Key point here is that
   all readers have the new key before anyone writes with it.
1. The list of DEKs being used by the master is updated in memory so that the
   new key is is the current write key.
1. The list of DEKs being used by the master is updated in etcd so that the new
   key is the current write key and is available to all masters. It doesn't
   matter if there's some masters using the new key and some using the old key,
   since we know all masters can read with the new key. Eventually all masters
   will be writing with the new key.
1. All secrets are re-encrypted with the new key.

After N rotation calls:

1. A new encryption key (unencrypted DEK) is created.
1. Encrypt DEK with KEK[N+1]

Each rotation generates a new KEK and DEK. Two DEKs will be in-use temporarily
during rotation, but only one at steady-state.

Pros:

 - Most closely matches the pattern that will be used for integrating with
   external encryption systems. Hashicorp Vault, Amazon KMS, Google KMS and HSM
   would eventually serve the purpose of KEK storage rather than local disk.

Cons:

 - End state is still KEKs on disk on the master. This is equivalent to the much
   simpler list of keys on disk in terms of key management and security.
   Complexity is much higher.
 - Coordination of keys between a deployment with multiple masters will require
   manually generating and providing a key in the key file then calling rotate
   to have the config re-read. Same as keys on disk.

### Option 4: Protocol for KEK agreement between masters

TODO: write a proposal for coordinating KEK agreement among multiple masters and
having the KEK be either user supplied or backed by external store.


## External providers

It should be easy for the user to substitute a default encryption provider for
one of the following:

* A local HSM implementation that retrieves the keys from the secure enclave
  prior to reusing the AES-GCM implementation (initialization of keys only)
* Exchanging a local temporary token for the actual decryption tokens from a
  networked secret vault
* Decrypting the AES-256 keys from disk using asymmetric encryption combined
  with a user input password
* Sending the data over the network to a key management system for encryption
  and decryption (Google KMS, Amazon KMS, Hashicorp Vault w/ Transit backend)

### Backwards Compatibility

Once a user encrypts any resource in etcd, they are locked to that Kubernetes
version and higher unless they choose to manually decrypt that resource in etcd.
This will be discouraged. It will be highly recommended that users discern if
their Kubernetes cluster is on a stable version before enabling encryption.

### Performance

Introducing even a relatively well tuned AES-GCM implementation is likely to
have performance implications for Kubernetes.  Fortunately, existing
optimizations occur above the storage layer and so the highest penalty will be
incurred on writes when secrets are created or updated.  In multi-tenant Kube
clusters secrets tend to have the highest load factor (there are 20-40 resources
types per namespace, but most resources only have 1 instance where secrets might
have 3-9 instances across 10k namespaces).  Writes are uncommon, creates usually
happen only when a namespace is created, and reads are somewhat common.

### Actionable Items / Milestones

* [p0] Add ValueTransformer to storage (Done in [#41939](https://github.com/kubernetes/kubernetes/pull/41939))
* [p0] Create a default implementation of AES-GCM interface (Done in [#41939](https://github.com/kubernetes/kubernetes/pull/41939))
* [p0] Add encryption flags on kube-apiserver and key rotation API
* [p1] Add kubectl command to call /rotate endpoint
* [p1] Audit of default implementation for safety and security
* [p2] E2E and performance testing
* [p2] Documentation and users guide
* [p2] Read cache layer if encrypting/decrypting Secrets adds too much load on kube-apiserver


## Alternative Considered: Encrypting the entire etcd database

It should be easy for the user to substitute a default encryption provider for
one of the following:

Rather than encrypting individual resources inside the etcd database, another
approach is to encrypt the entire database.

Pros:

 - Removes the complexity of deciding which types of things should be encrypted
   in the database.
 - Protects any other sensitive information that might be exposed if etcd
   backups are made public accidentally or one of the other desribed attacks
   occurs.

Cons:

 - Unknown, but likely significant performance impact. If it isn't fast enough
   you don't get to fall back on only encrypting the really important stuff.
   As a counter argument: Docker [implemented their encryption at this
   layer](https://docs.docker.com/engine/swarm/swarm_manager_locking/) and have
   been happy with the performance.

