# Kubectl Conventions

**Table of Contents**

- [Kubectl Conventions](#kubectl-conventions)
  - [Principles](#principles)
  - [Command conventions](#command-conventions)
    - [Create commands](#create-commands)
    - [Plugins](#plugins)
  - [Flag conventions](#flag-conventions)
  - [Output conventions](#output-conventions)
  - [Documentation conventions](#documentation-conventions)
  - [kubectl client conventions](#kubectl-client-conventions)
  - [Command implementation conventions](#command-implementation-conventions)
  - [Exit code conventions](#exit-code-conventions)
  - [Verification](#verification)


## Principles

* Strive for consistency across commands.

* Explicit should always override implicit.

  * User preferences from `kuberc` should override default values.

  * Environment variables should override default values.

  * Command-line flags should override default values, `kuberc` preferences,
    and environment variables.

  * The namespace set in a resource (e.g., in a file passed via `-f`) overrides
    the default namespace from kubeconfig. An explicit `--namespace` that
    conflicts with the namespace set in a resource is an error, rather than
    an override, to guard against accidentally operating outside the intended
    namespace.

* Most kubectl commands should be able to operate in bulk on resources of mixed types.

* Kubectl should not make any decisions based on its own or the server's release version
  string. Instead, API discovery and/or OpenAPI should be used to determine available features.

* We currently only guarantee support for [one release of version skew](https://kubernetes.io/releases/version-skew-policy/#kubectl),
  but we strive to make old releases of kubectl continue to work with newer servers
  in compliance with our API compatibility guarantees. This means, for instance,
  that kubectl should not fully parse objects returned by the server into full Go
  types and then re-encode them, since that would drop newly added fields.

* General-purpose kubectl commands (e.g., `get`, `delete`, `create`, `replace`,
  `patch`, `apply`) should work for all resource types, even those not present
  when that release of kubectl was built, such as APIs added in newer releases,
  aggregated APIs, and custom resources.

* While functionality may be added to kubectl out of expedience, commonly needed
  functionality should be provided by the server to make it easily accessible
  to all API clients. Examples of functionality that moved server-side: `get` table
  output, `apply` (server-side apply), dry-run, field validation, and resource
  categories.

* Remaining non-trivial functionality in kubectl should be made available to other
  clients via libraries. Code lives in `staging/src/k8s.io/kubectl` (published as
  `k8s.io/kubectl`) and the reusable CLI building blocks (config flags, print flags,
  resource builder, printers) live in `staging/src/k8s.io/cli-runtime` (published
  as `k8s.io/cli-runtime`).

* Experimental client-side behavior is gated behind `KUBECTL_*` environment variables
  (see `FeatureGate` in `staging/src/k8s.io/kubectl/pkg/cmd/util/helpers.go`,
  e.g., `KUBECTL_KUBERC`, `KUBECTL_APPLYSET`). Gates that are disabled by default
  (usually at alpha stage) are turned on by setting the variable to `true`; gates
  that are enabled by default (usually at beta stage) can be turned off by setting
  the variable to `false`.


## Command conventions

* Command names are all lowercase, and hyphenated if multiple words.

* Use `kubectl <VERB> <NOUNs>` for commands that apply to multiple resource types.

* Commands should not have built-in aliases. Users who want aliases can
  define them in `kuberc`. The exception is subcommands named after a resource
  type (e.g., `kubectl create deployment`, `kubectl top pod`), which accept that
  resource's other names as aliases (`deploy`, `pods`, `po`).

* `<NOUNs>` may be specified as `TYPE name1 name2` or `TYPE/name1 TYPE/name2`.

* Resource types are all lowercase, with no hyphens; both singular and plural
  forms are accepted. Types may be fully qualified as `resource.version.group`
  or `resource.group` to disambiguate.

* `<NOUNs>` may also be specified by one or more file arguments: `-f file1 -f file2
  ...`, a directory (`-f dir/`, optionally with `-R`), a URL, or a kustomization
  directory (`-k dir/`).

* Resource types may have 2- or 3-letter short names. Short names are served by
  the API server through discovery, not hard-coded in kubectl.

* Business logic should be decoupled from the command framework, so that it can
  be reused independently of `kubectl`, `cobra` library, etc.

  * Ideally, commonly needed functionality would be implemented server-side in
    order to avoid problems typical of "fat" clients and to make it readily
    available to non-Go clients.

* A command group (e.g., `kubectl config`, `kubectl set`, `kubectl rollout`,
  `kubectl create`) may be used to group related non-standard commands, such as
  object construction, mutations, and computations.

* Top-level commands are registered in `staging/src/k8s.io/kubectl/pkg/cmd/cmd.go`.
  Most belong to one of the `templates.CommandGroups`, shown in help as
  "Basic Commands (Beginner)", "Basic Commands (Intermediate)", "Deploy Commands",
  "Cluster Management Commands", "Troubleshooting and Debugging Commands",
  "Advanced Commands", and "Settings Commands". A few (`config`, `plugin`,
  `version`, `api-resources`, `api-versions`, `options`, `kuberc`) are added
  outside the groups and listed under "Other Commands" in help; installed plugins
  get their own help group. Commands that are not ready for general use go under
  `kubectl alpha` (`staging/src/k8s.io/kubectl/pkg/cmd/alpha.go`), which is hidden
  when empty.

### Create commands

`kubectl create <resource>` commands fill the gap between "I want to try
Kubernetes, but I don't know or care what gets created" (`kubectl run`, which
creates a single pod) and "I want to create exactly this" (write YAML and run
`kubectl create -f`). They provide an easy way to create a valid object without
having to know the vagaries of particular kinds, nested fields, and object key
typos that are ignored by the YAML/JSON parser. Because editing an already
created object is easier than authoring one from scratch, these commands only
need to have enough parameters to create a valid object and set common
immutable fields. They should default as much as is reasonably possible. Once
that valid object is created, it can be further manipulated using `kubectl
edit` or `kubectl set` commands.

`kubectl create <resource> <special-case>` commands help in cases where you need
to perform non-trivial configuration generation/transformation tailored for a
common use case. `kubectl create secret` is a good example: there's a `generic`
flavor with keys mapping to files, then there's a `docker-registry` flavor that
is tailored for creating an image pull secret, and there's a `tls` flavor for
creating TLS secrets. `kubectl create service` follows the same pattern with
`clusterip`, `nodeport`, `loadbalancer`, and `externalname`. These are separate
commands so that each gets its own flags and help tailored to its use.

### Plugins

* Any executable on `PATH` named `kubectl-<name>` is invoked for `kubectl <name>`,
  with dashes in the filename mapping to subcommands (`kubectl-foo-bar` -> `kubectl foo bar`)
  and underscores mapping to dashes in the command name.

* Plugins can never override a built-in command. The only built-in command that
  accepts plugin subcommands is `create` (`kubectl-create-foo` -> `kubectl create foo`),
  and only for subcommands that don't exist as built-ins.

* Functionality that is useful to a subset of users, or that is still being
  iterated on, is a good candidate for a plugin rather than a built-in command.


## Flag conventions

* Flags are all lowercase, with words separated by hyphens. This is enforced by
  `hack/verify-cli-conventions.sh`. Flags using `_` are normalized to `-` with a
  warning.

* Flag names and short (single-character) flags should have the same meaning across all commands.

* Flag descriptions should start with an uppercase letter. Full sentences
  should end with a period; most existing flags do, but a short phrase without
  one is acceptable.

* Command-line flags corresponding to API fields should accept API enums
  exactly (e.g., `--restart=Always`).

* Do not reuse flags for different semantic purposes, and do not use different
  flag names for the same semantic purpose. Check `staging/src/k8s.io/kubectl/pkg/cmd/util/helpers.go`
  (`Add*Flag*` helpers) and `staging/src/k8s.io/cli-runtime/pkg/genericclioptions`
  (`*Flags` structs) for an existing flag before adding a new one, and use the
  shared helper rather than redefining the flag.

* Use short flags sparingly, only for the most frequently used options, prefer
  lowercase over uppercase for the most common cases, try to stick to well-known
  conventions for UNIX commands, where they exist, and update this list when
  adding new short flags.

  * `-A`: All namespaces
  * `-c`: Container
    * also used for `--containers` in `set env` and `set resources`
  * `-e`: Environment variable, in `set env`
  * `-f`: Resource file
    * also used for `--follow` in `logs`, but should be deprecated in favor of `-F`
  * `-h`: Help
  * `-i`: Attach stdin
    * also used for `--interactive` in `delete`
  * `-k`: Kustomization directory
  * `-l`: Label selector
    * also used for `--labels` in `expose` and `run`, but should be deprecated
  * `-L`: Label columns
  * `-n`: Namespace scope
  * `-o`: Output format
  * `-p`: Previous, in `logs`
    * also used for `--patch` in `patch`, but should be deprecated
    * also used for `--port` in `proxy`, but should be deprecated
  * `-P`: Static file prefix (`--www-prefix`) in `proxy`, but should be deprecated
  * `-q`: Quiet
  * `-r`: Replicas, in `create deployment`
  * `-R`: Recursive
  * `-s`: API server address
  * `-t`: Allocate TTY
  * `-u`: Unix socket, in `proxy`
  * `-v`: Verbose logging level
  * `-w`: Watch
    * also used for `--www` in `proxy`, but should be deprecated

* `--dry-run=none|client|server`: Don't modify the live state. `client` only
  prints the object that would be sent; `server` submits the request without
  persisting it. All mutations should support it via `cmdutil.AddDryRunFlag`
  and `cmdutil.GetDryRunStrategy`.

* `--local`: Don't contact the server; perform only local reads, transformations,
  generation, etc., and display the output.

* `--field-manager`: Name of the manager used to track field ownership, added
  via `cmdutil.AddFieldManagerFlagVar`. Defaults to `kubectl-<verb>` (e.g.,
  `kubectl-create`, `kubectl-set`, `kubectl-rollout`). All mutations should support it.

* `--validate=strict|warn|ignore`: Schema validation of the input, added via
  `cmdutil.AddValidateFlags`. `true` and `false` are accepted as aliases for
  `strict` and `ignore`. Validation is performed server-side when the server
  supports field validation; otherwise, kubectl falls back to client-side
  validation for `strict`. Defaults to `strict`.

* `--record` is deprecated; don't add it to new commands.


## Output conventions

* By default, output is intended for humans rather than programs.

  * However, affordances are made for simple parsing of `get` output.

* stdout carries the result of the command; everything else (errors, warnings,
  and status messages such as `No resources found in default namespace.`)
  goes to stderr, so that stdout stays safe to pipe into other programs. Warnings
  returned by the server are printed to stderr (deduplicated), and
  `--warnings-as-errors` turns them into a non-zero exit code.

* `get` commands should output one row per resource, and one resource per row.

  * Columns for built-in types are defined server-side as `Table` column
    definitions (`pkg/printers/internalversion` in the main repo); custom resources
    use `additionalPrinterColumns`. `kubectl` does not hard-code per-type columns.

  * New column titles and values should not contain spaces, so that lines can be
    split into fields with `cut`, `awk`, etc. Instead, use `-` as the word
    separator.

  * By default, `get` output should fit within about 80 columns.

    * `-o wide` may be used to display additional columns.

  * The first column should be the resource name, titled `NAME`.

  * `NAMESPACE` should be displayed as the first column when `--all-namespaces`
    is specified.

  * The last default column should be time since creation, titled `AGE`.

  * `-L <key>` appends a column containing the value of the label with key `key`,
    left empty if not present. The column title is the uppercased key without
    its prefix (`-L app.kubernetes.io/name` produces `NAME`).

  * The `json`, `yaml`, `kyaml`, `name`, Go template (`go-template`, `go-template-file`),
    and jsonpath (`jsonpath`, `jsonpath-file`, `jsonpath-as-json`) formats should
    be supported and encouraged for subsequent processing. Commands get these
    through `genericclioptions.PrintFlags` rather than wiring printers by hand.
    `get` additionally supports `wide`, custom columns (`custom-columns`,
    `custom-columns-file`), and `-L` through its own `PrintFlags`
    (`staging/src/k8s.io/kubectl/pkg/cmd/get/get_flags.go`).

* `describe` commands may output on multiple lines and may include information
  from related resources, such as events. Describe should include information
  from related resources that a typical user may need to know. If a user would
  always run "describe resource1" and then immediately want to run a "get type2"
  or "describe resource2", consider including that info. Examples: persistent
  volume claims for pods that reference claims, events for most resources, and
  nodes and the pods scheduled on them. When fetching related resources, a
  targeted field selector should be used instead of client-side filtering of
  related resources.

* In `describe` output, fields that can be explicitly unset (booleans, integers,
  structs) should show `<unset>`. Likewise, arrays should show `<none>`. Finally,
  `<unknown>` should be used where an unrecognized value was specified.
  Server-side `get` columns follow the same placeholders and add `<pending>` for
  a load balancer address that has not been assigned yet.

* Mutations should output `TYPE/name verbed` by default, where `TYPE` is the
  lowercase singular kind, qualified with the group for non-core types (e.g.,
  `pod/nginx created`, `deployment.apps/nginx scaled`). `-o name` prints only
  `TYPE/name`, which other commands accept as input.


## Documentation conventions

* Commands are documented using `cobra`; docs are then auto-generated by
  `hack/update-generated-docs.sh`.

  * `Use` should contain a short usage string for the most common use case(s), not
    an exhaustive specification. Set `DisableFlagsInUseLine: true` and list the
    relevant flags in `Use` explicitly.

  * `Short` should contain a one-line explanation of what the command does.

    * Short descriptions should start with an uppercase letter and not end
      with a period.

    * Short descriptions should (if possible) start with a verb in the
      imperative mood.

  * `Long` may contain multiple lines, including additional information about
    input, output, commonly used flags, etc.

    * Long descriptions should use proper grammar, start with an uppercase
      letter, and end sentences with a period.

    * `Long` must be wrapped in `templates.LongDesc(...)` (enforced by
      `hack/verify-cli-conventions.sh`).

  * `Example` should contain usage examples covering the common cases.

    * `Example` must be wrapped in `templates.Examples(...)` (enforced by
      `hack/verify-cli-conventions.sh`).

    * A comment should precede each example command. Comments use `#`, not
      `//` (enforced), and should start with an uppercase letter.

    * Command examples should not include a `$` prefix.

  * `Short`, `Long`, and `Example` strings should be wrapped in `i18n.T(...)` so
    they can be translated (`hack/update-translations.sh`).

* Use `FILENAME` for filenames.

* Use `TYPE` for the particular flavor of resource type accepted by kubectl,
  rather than `RESOURCE` or `KIND`.

* Use `NAME` for resource names.


## kubectl client conventions

The kubectl `Factory` (`staging/src/k8s.io/kubectl/pkg/cmd/util/factory.go`) is
an interface that provides access to clients (`DynamicClient`, `KubernetesClientSet`,
`RESTClient`, `ClientForMapping`, `UnstructuredClientForMapping`), the resource
builder (`NewBuilder`), schema validation (`Validator`), and OpenAPI v2/v3 schemas.
It embeds `genericclioptions.RESTClientGetter`, which provides the REST config,
discovery client, REST mapper, and kubeconfig loader.

Commands should depend on the narrowest interface that works. Commands that only
need config, discovery, or a REST mapper should accept a `genericclioptions.RESTClientGetter`
instead of the full `Factory` (e.g., `events`, `wait`, `debug`).


## Command implementation conventions

For every command there should be a `NewCmd<CommandName>` function that creates
the command and returns a pointer to a `cobra.Command`, which can later be added
to other parent commands to compose the structure tree. It takes a `genericclioptions.RESTClientGetter`
and a `genericiooptions.IOStreams`. Commands never write to `os.Stdout`/`os.Stderr`
directly.

There should be a `<CommandName>Flags` struct with a field for every flag and
argument declared by the command, created by a `New<CommandName>Flags(restClientGetter, streams)`
constructor. The `<CommandName>Flags` struct should define two member methods:

* `AddFlags`: responsible for wiring flags using the `cobra` library.

* `ToOptions`: responsible for translating flags into a `<CommandName>Options`
  struct and setting all other required values.

The `<CommandName>Options` struct contains a field for every flag and argument
declared by the command, and any other field required for the command to run.
This makes tests and mocking easier. The `<CommandName>Options` struct ideally
exposes two methods:

* `Validate`: performs validation on the struct fields and returns appropriate
  errors.

* `Run`: runs the actual logic of the command, assuming the struct is complete
  and has been validated. It takes a `context.Context`, so that cancellation
  propagates from `cmd.Context()`. Some commands might use `Run<CommandName>`,
  but just `Run` is preferred.

The `cobra.Command`'s `Run` function calls `ToOptions`, `Validate`, and `Run` in that order.

Sample command skeleton:

```go
var (
	mineLong = templates.LongDesc(i18n.T(`
		Mine which is described here
		with lots of details.`))

	mineExample = templates.Examples(i18n.T(`
		# Run my command's first action
		kubectl mine first_action

		# Run my command's second action on latest stuff
		kubectl mine second_action --latest`))
)

// MineFlags contains all the flags and arguments for the mine CLI command.
type MineFlags struct {
	RESTClientGetter genericclioptions.RESTClientGetter

	Latest bool

	genericiooptions.IOStreams
}

// NewMineFlags returns a default MineFlags.
func NewMineFlags(restClientGetter genericclioptions.RESTClientGetter, streams genericiooptions.IOStreams) *MineFlags {
	return &MineFlags{
		RESTClientGetter: restClientGetter,
		IOStreams:        streams,
	}
}

// NewCmdMine implements the kubectl mine command.
func NewCmdMine(restClientGetter genericclioptions.RESTClientGetter, streams genericiooptions.IOStreams) *cobra.Command {
	flags := NewMineFlags(restClientGetter, streams)

	cmd := &cobra.Command{
		Use:                   "mine ACTION [--latest]",
		DisableFlagsInUseLine: true,
		Short:                 i18n.T("Run my command"),
		Long:                  mineLong,
		Example:               mineExample,
		Run: func(cmd *cobra.Command, args []string) {
			o, err := flags.ToOptions(args)
			cmdutil.CheckErr(err)
			cmdutil.CheckErr(o.Validate())
			cmdutil.CheckErr(o.Run(cmd.Context()))
		},
	}

	flags.AddFlags(cmd)

	return cmd
}

// AddFlags registers flags for the mine command.
func (flags *MineFlags) AddFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&flags.Latest, "latest", flags.Latest, "Use latest stuff.")
}

// ToOptions converts from CLI inputs to runtime inputs.
func (flags *MineFlags) ToOptions(args []string) (*MineOptions, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("exactly one ACTION is required, got %d", len(args))
	}

	namespace, _, err := flags.RESTClientGetter.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return nil, err
	}

	return &MineOptions{
		Action:    args[0],
		Latest:    flags.Latest,
		Namespace: namespace,
		IOStreams: flags.IOStreams,
	}, nil
}

// MineOptions contains all the options for running the mine CLI command.
type MineOptions struct {
	Action    string
	Latest    bool
	Namespace string

	genericiooptions.IOStreams
}

// Validate validates all the required options for mine.
func (o *MineOptions) Validate() error {
	return nil
}

// Run implements all the necessary functionality for mine.
func (o *MineOptions) Run(ctx context.Context) error {
	return nil
}
```

Most existing commands predate this pattern and have no `<CommandName>Flags`, only
`<CommandName>Options` with three methods: `Complete` (which acts similarly to
`ToOptions`), `Validate`, and `Run`. The downside of this approach is that it
tightly couples commands to the `cobra` library. New commands should use
`<CommandName>Flags`; `wait` and `events` are good references.


## Exit code conventions

In general, an exit code of `0` means success and any non-zero exit code means failure.

| Exit code | Meaning |
| :---      | :---    |
| 0         | Success |
| 1         | General error |
| other     | Propagated from a remote process, e.g., the container exit code for `kubectl exec` and attached `kubectl run` |

Commands may define their own exit code contract when it mirrors a well-known
tool; such commands must document it in `Long`. `kubectl diff` follows `diff(1)`:
`0` when no differences were found, `1` when differences were found, and `>1`
when kubectl or diff failed. It uses `cmdutil.CheckDiffErr` so that kubectl errors
(including flag parsing errors) never exit with `1`.


## Verification

* `hack/verify-cli-conventions.sh` checks that every command's `Long` and `Example`
  are normalized through `templates`, that examples use `#` comments, that long
  flag names contain only lowercase letters and dashes, and that no flags are
  registered on the global `flag.CommandLine`.

* `hack/verify-generated-docs.sh` checks that generated docs are up to date.

* Unit tests live next to the command
  (`staging/src/k8s.io/kubectl/pkg/cmd/<name>/<name>_test.go`).

* Integration tests for kubectl live in `test/cmd` (shell-based, running against
  an API server without a kubelet).

* Full end-to-end tests for kubectl live in `test/e2e/kubectl`.
