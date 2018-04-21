# Encryption at Rest - Integration with Google KMS

## Abstract

The scope of this proposal is to cover the implementation of an encryption
provider for Google KMS that will encrypt data as it is stored to etcd. This
proposal is specifically about the Google KMS integration. The general
description of the Encryption at Rest feature is covered in [PR
#607](https://github.com/kubernetes/community/pull/607/files).

## High Level Design

The design described in this documentation is the implementation of an external
provider as described in [PR #607](https://github.com/kubernetes/community/pull/607/files).

See https://github.com/kubernetes/kubernetes/pull/48574 for alpha implementation
details.

The KMS integration will make remote calls to KMS to encrypt and decrypt data
instead of using local encryption as detailed in [PR #607](https://github.com/kubernetes/community/pull/607/files).

In order to:
- make key rotation easier
- decrease the amount of data sent over the network
- avoid hitting the upper limit of data that can be encrypted by a call to
  Google KMS

the implementation will use envelope encryption. A Data Encryption Key (DEK)
will be generated locally and used to encrypt the data destined for etcd. The
DEK will be encrypted with a call to KMS, and the resulting cypher text will be
prepended to the data destined for etcd. Each write will generate a new DEK.

When reading, DEKs will be cached in the api server to facilitate faster reads.
The first time a piece of data is read, the DEK will not be available locally,
so the DEK will be sent to KMS and the resulting plain text of the DEK will be
stored in memory.

