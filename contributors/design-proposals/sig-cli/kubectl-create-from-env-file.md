# Kubectl create configmap/secret --env-file

## Goals

Allow a Docker environment file (.env) to populate an entire `ConfigMap` or `Secret`.
The populated `ConfigMap` or  `Secret` can be referenced by a pod to load all
the data contained within.

## Design

The `create configmap` subcommand would add a new option called
`--from-env-file`. The option will accept a single file. The option may not be
used in conjunction with `--from-file` or `--from-literal`.

The `create secret generic` subcommand would add a new option called
`--from-env-file`. The option will accept a single file. The option may not be
used in conjunction with `--from-file` or `--from-literal`.

### Environment file specification

An environment file consists of lines to be in VAR=VAL format. Lines beginning
with # (i.e. comments) are ignored, as are blank lines. Any whitespace in
front of the VAR is removed. VAR must be a valid C_IDENTIFIER.  If the line
consists of just VAR, then the VAL will be given a value from the current
environment.

Any ill-formed line will be flagged as an error and will prevent the
`ConfigMap` or `Secret` from being created.

[Docker's environment file processing](https://github.com/moby/moby/blob/master/opts/env.go)

## Examples

```
$ cat game.env
enemies=aliens
lives=3
enemies_cheat=true
enemies_cheat_level=noGoodRotten
secret_code_passphrase=UUDDLRLRBABAS
secret_code_allowed=true
secret_code_lives=30
```

Create configmap from an env file:
```
kubectl create configmap game-config --from-env-file=./game.env
```

The populated configmap would look like:
```
$ kubectl get configmaps game-config -o yaml

apiVersion: v1
data:
  enemies: aliens
  lives: 3
  enemies_cheat: true
  enemies_cheat_level: noGoodRotten
  secret_code_passphrase: UUDDLRLRBABAS
  secret_code_allowed: true
  secret_code_lives: 30
```

Create secret from an env file:
```
kubectl create secret generic game-config --from-env-file=./game.env
```

The populated secret would look like:
```
$ kubectl get secret game-config -o yaml

apiVersion: v1
type: Opaque
data:
  enemies: YWxpZW5z
  enemies_cheat: dHJ1ZQ==
  enemies_cheat_level: bm9Hb29kUm90dGVu
  lives: Mw==
  secret_code_allowed: dHJ1ZQ==
  secret_code_lives: MzA=
  secret_code_passphrase: VVVERExSTFJCQUJBUw==
```
