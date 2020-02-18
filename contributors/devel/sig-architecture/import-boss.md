# Using import-boss to enforce package import restrictions

This document shows a way of enforcing package import restrictions using the `import-boss` tool.

## Background

The import-boss check runs against all pull requests submitted against the k/k repository. There are currently a good
number of import restriction files, called `.import-restrictions` in the repository. The import restrictions files are
specified in YAML format.

## Rules

Each import restrictions files contains rules. A Rule consists of three parts:

1. A `SelectorRegexp`, to select the import paths that the rule applies to.
2. A list of `AllowedPrefixes`.
3. A list of `ForbiddenPrefixes`.

All imports of a package are checked against each "rule" in the found restriction files, climbing up the directory tree
until the import matches one of the rules. An import matches a rule of a matching selector if it matches at least one
allowed prefix, but no forbidden prefix. If the import matches any of the rules, it is accepted.

## Inverse Rules

Reverse import rules help to define which packages in [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes)
can import certain other packages, defined inside the latter. This allows to have fine-grained import restrictions for
"private packages" where we don't want to spread use inside of [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes).

All incoming imports of the package are checked against each "inverse rule" in the found restriction files, climbing up
the directory tree until the import matches one of the rules. If the import matches any of the rules, it is accepted.

Inverse rules also have a boolean `Transitive` option. When this option is true, the import rule is also applied to transitive
imports.

## Example file

An example `.import-restrictions` file looks like this:

```yaml
rules:
  - selectorRegexp: k8s[.]io
    allowedPrefixes:
      - k8s.io/gengo/examples
      - k8s.io/kubernetes/third_party
    forbiddenPrefixes:
      - k8s.io/kubernetes/pkg/third_party/deprecated
  - selectorRegexp: ^unsafe$
    forbiddenPrefixes:
      - ""
inverseRules:
  - selectorRegexp: k8s[.]io
    allowedPrefixes:
      - k8s.io/same-repo
      - k8s.io/kubernetes/pkg/legacy
    forbiddenPrefixes:
      - k8s.io/kubernetes/pkg/legacy/subpkg
  - selectorRegexp: k8s[.]io
    transitive: true
    forbiddenPrefixes:
      - k8s.io/kubernetes/cmd/kubelet
      - k8s.io/kubernetes/cmd/kubectl
```

Note the second (non-inverse) rule explicitly matches the unsafe package, and forbids it ("" is a prefix of everything).

Also note that the second InverseRule is transitive, the first only applies to direct imports.


[import-boss](https://github.com/kubernetes/gengo/tree/master/examples/import-boss)
