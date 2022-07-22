This copy of [Versioned Component Configuration Files](vccf-proposal)
was automatically converted from Google Docs to Markdown
so that it could be included in the community archive.
Several interesting comment threads remain unexported but
available in the Google Doc. You can gain access to the
doc by joining the `dev@kubernetes.io` mailing list
(previously `kubernetes-dev@googlegroups.com`). Please
do not request access via the Google Docs UI, as this
spams the owners with access requests. Thank you :).

[vccf-proposal]: https://docs.google.com/document/d/1FdaEJUEh091qf5B98HM6_8MS764iXrxxigNIdwHYW9c/


# Versioned Component Configuration Files


## How Kubelet learned to stop using flags and love versioned config. How your component can, too.

**Shared publicly!**

**Author:** mtaufen@google.com

**Last Updated:** March 28, 2018

**Self Link:** https://goo.gl/GM8KyH 

# Background

A long time ago (but still in this galaxy), someone had the bright idea that we could avoid a lot of the pain of deploying and managing command-line flags for each core cluster component (kubelet, kube-proxy, scheduler, etc.) by switching to Kubernetes-style versioned configuration files. This effort became known to the community as _component configuration_, or simply _componentconfig_. At its origin, a consistent philosophy for _what_ componentconfig should look like did not exist.

Last year, mikedanese@ did a great job of compiling the ideas behind componentconfig (ideas Mike, other folks, and I were discussing in several GitHub threads) into a single [document](https://docs.google.com/document/d/1arP4T9Qkp2SovlJZ_y790sBeiWXDO6SG10pZ_UUU-Lc), in the hope that we could provide standard guidelines and improve consistency across the project. Mike's document catalyzed our push to try componentconfig in a few components, and here we are almost a year later.

As of Kubernetes v1.10, the Kubelet is firmly on its way to migrating from flags to versioned configuration files. It can consume a beta-versioned config file and many flags are now deprecated and pending removal in favor of this file. Many remaining flags will be replaced by the file over time. Additionally, the kube-proxy component is very close to having a beta-versioned config file of its own.

This document restates the motivation and records lessons from OSS work over the past year:



*   a brief review of why we want versioned config files for all core cluster components
*   the _ideal_ state of a component's configuration API
*   how to migrate an existing component to versioned configuration files (Kubelet example)
*   remaining work


# Why versioned config files?

The short answer is that flags are nonstandard interfaces with weak stability guarantees. They are confusing and hard to deploy, and this is the opposite of what Kubernetes should be.

Command-line flags present a number of problems:



*   Flags are a public API, but are not versioned and cannot be versioned separately from the binary: 
    *   For _core_ components the binary version is coupled to the Kubernetes release version. We use semantic versioning for our binaries, but can't bump a major version unless Kubernetes does:
        *   We _shouldn't_ ever **fix bad defaults** for existing flag names without bumping the major version of a binary. In reality, we get around this by giving advance warning and technically breaking semantic versioning of the binary.
        *   We _shouldn't_ ever **remove a flag** without bumping the major version of a binary. In reality we use a [flag deprecation policy](https://kubernetes.io/docs/reference/deprecation-policy/#deprecating-a-flag-or-cli) that allows us to technically break semantic versioning of the binary as long as we give advance warning.
    *   We incrementally deprecate individual parameters over time, instead of guaranteeing a consistent set of parameters for the life of an API version. This confuses users and results in a less stable API.
    *   We can't typically deploy flag-based configuration independently from a binary version upgrade, because the compatibility of the interface is so tightly coupled.
*   Values are often re-configured, which precipitates additional tools to parameterize and write configuration for system-specific process management agents (e.g. systemd). We can eliminate the dependency on parameterization tools if process manager config is static; e.g. the configured command line just needs to reference a file in a fixed location.
*   Developers inevitably embed structured data in strings and invent one-off parsers to process their flags. This invites bugs.
*   mikedanese@ outlined more issues in his [document](https://docs.google.com/document/d/1arP4T9Qkp2SovlJZ_y790sBeiWXDO6SG10pZ_UUU-Lc).

Core goals of componentconfig include:



*   Standardize the configuration approach for all core cluster components.
*   Enable dynamic configuration deployment mechanisms.

Conveniently, Kubernetes has similar goals:



*   Standardize the configuration approach for cluster infrastructure.
*   Enable dynamic deployment mechanisms.

Kubernetes had already paved the way: It has what we need to version our configuration interfaces, decouple configuration changes from binary changes, represent configuration in a structured format, and deploy configuration in a dynamic environment.



*   Versioning was accomplished via the API machinery's group/version mechanism.
*   Adhering to the same API guarantees as the core Kubernetes APIs provides a stable configuration surface, and allows us to decouple the configuration interface from releases that support the same API version.
*   Kubernetes API objects are defined as Go structs, which means we don't have to parse strings to deal with structured config.
*   Kubernetes has deployment mechanisms (ConfigMap volume source) that work well for pushing new versions of configuration files into production.
*   There is no requirement to restart a process when you change a file, unlike flags.

**All core Kubernetes components should eventually consume their configuration via versioned configuration files, instead of command-line flags.**


# tl;dr: What should a component do?

**This is the _ideal_ command-line API for every core cluster component:**

```console
$ component --config=path
```

The component exposes _only one_ flag on its command line. This flag provides the file path to a config file with a versioned format. All other relevant configuration information is referenced via this file. 

_One, stable flag where everything else is versioned config_ is the ideal API recommended by componentconfig. If you are creating a new component from scratch, begin and end with this API. 

For several reasons discussed in the next section, the migration from flags to versioned config is a serious journey for most existing components, and these may want or even need a couple more flags at the end of the day.

In general, every core cluster component should:



*   Maintain a distinct Kubernetes API group called {component}.config.k8s.io, which contains versioned sets of config objects - primarily a {Component}Configuration struct in each version. This struct, serialized to disk by the API machinery, is the file format for configuration.
*   Ensure {component}.config.k8s.io adheres to the standard Kubernetes [API deprecation policy](https://kubernetes.io/docs/reference/deprecation-policy/#deprecating-parts-of-the-api), [API conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md), and [API changes policy](https://github.com/kubernetes/community/blob/master/contributors/devel/api_changes.md).
*   Expose a flag named --config, which accepts a path to a file that contains a serialized {Component}Configuration struct.
*   Use the Kubernetes API machinery to deserialize the config file data, apply defaults, and convert to an internal version for runtime use.
*   Validate the internal version prior to using it. If validation fails, refuse to run with the specified configuration.
*   Ensure third-party libraries aren't leaking flags.

We discuss how to migrate a component from flags to versioned config files in the next section.


# In-depth: How to migrate a component from flags to a versioned config file


## Take back control of the command-line API

Our goal is to _decrease_ flag usage in favor of versioned config files. It will help to decrease the growth rate of the component's flag API. There are at least two sources of this growth:



*   PRs that directly extend the flag API.
*   Adding or updating third-party libraries.

Whoever owns the componentconfig effort for a given component should be in-the-loop on PRs that add new flags. This person has a strong interest in saying "no" to new flags, because it increases the number of things they have to carefully migrate to the new versioned config file API. When new flags really prove necessary, this person still has a strong interest in ensuring they will be compatible when migrated.

The second case is an artifact of many libraries registering flags globally (global flag sets are provided by both the _flag_ and _pflag_ libraries). Since most components just parse the global flag sets by default, they tend to accumulate the flags from these libraries. The libraries are impolite, and the components are typically too trusting. Each component should be more cautious by:



*   constructing its own, isolated flag set
*   explicitly registering necessary flags from third-party libraries into this flag set
*   parsing _only_ the flags in this flag set

You can find the example of how the Kubelet took back control of its flag set in [Explicit kubelet flags](https://github.com/kubernetes/kubernetes/pull/57613) (see also the follow-up PR, [#58095](https://github.com/kubernetes/kubernetes/pull/58095)).

Many components indiscriminately add the global flag set to their primary flag set via `pflag.CommandLine.AddGoFlagSet(flag.CommandLine)`. Further, most components delegate their flag parsing, help text generation, etc. to [Cobra](https://github.com/spf13/cobra). Cobra implicitly adds flag.CommandLine in several cases, which unfortunately gets in the way of explicit control over the flag API.

In order for Cobra to parse flags for you, it has to be made aware of your flagset. This is achieved by registering flags to the command's flag set. Cobra will implicitly merge the global command lines with this flag set when it parses flags. Consider the following Go program, which creates a local, isolated flag set (as recommended above). The program also registers a global flag, which we _hope_ won't be parsed, because it is not explicitly registered with our local flag set. 


```go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const use = "testcmd"

var (
	globalFlagTarget string
	localFlagTarget  string
)

func init() {
	pflag.StringVar(&globalFlagTarget, "global-flag", globalFlagTarget, "globally-registered flag")
}

func NewLocalFlagSet() *pflag.FlagSet {
	fs := pflag.NewFlagSet(use, pflag.ContinueOnError)
	fs.StringVar(&localFlagTarget, "local-flag", localFlagTarget, "locally-registered flag")
	return fs
}

func NewTestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: use,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("globalFlagTarget: %q\n", globalFlagTarget)
			fmt.Printf("localFlagTarget: %q\n", localFlagTarget)
		},
	}
	cmd.Flags().AddFlagSet(NewLocalFlagSet())
	return cmd
}

func main() {
	cmd := NewTestCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
```


If we run the program, we see that this is not the case. The global flag is parsed by Cobra:


```console
$ testcmd --global-flag hello

globalFlagTarget: "hello"
localFlagTarget: ""
```


You can circumvent this by disabling Cobra's flag parsing. This, unfortunately, requires that you both parse flags and implement --help short-circuiting on your own. The next example attempts to extend the above to do so.


```go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const use = "testcmd"

var (
	globalFlagTarget string
	localFlagTarget  string
)

func init() {
	pflag.StringVar(&globalFlagTarget, "global-flag", globalFlagTarget, "globally-registered flag")
}

func NewLocalFlagSet() *pflag.FlagSet {
	fs := pflag.NewFlagSet(use, pflag.ContinueOnError)
	fs.StringVar(&localFlagTarget, "local-flag", localFlagTarget, "locally-registered flag")
	return fs
}

func NewTestCmd() *cobra.Command {
	localFlagSet := NewLocalFlagSet()
	cmd := &cobra.Command{
		Use:                use,
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			// parse our local flag set
			if err := localFlagSet.Parse(args); err != nil {
				cmd.Usage()
				fatal(err)
			}
			// --help
			help, err := localFlagSet.GetBool("help")
			if err != nil {
				fatal(fmt.Errorf(`"help" flag is non-bool, programmer error, please correct`))
			}
			if help {
				cmd.Help()
				return
			}
			// print the flag values
			fmt.Printf("globalFlagTarget: %q\n", globalFlagTarget)
			fmt.Printf("localFlagTarget: %q\n", localFlagTarget)
		},
	}
	localFlagSet.BoolP("help", "h", false, fmt.Sprintf("help for %s", cmd.Name()))
	// Cobra still needs to be aware of our flag set to generate usage and help text
	cmd.Flags().AddFlagSet(localFlagSet)
	return cmd
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func main() {
	cmd := NewTestCmd()
	if err := cmd.Execute(); err != nil {
		fatal(err)
	}
}
```


When we run the program and attempt to set the global flag, we see that the global flag is now rejected, but we also see that it is still included in the usage text!


```console
$ testcmd --global-flag hello
Usage:
  testcmd [flags]

Flags:
      --global-flag string   globally-registered flag
  -h, --help                 help for testcmd
      --local-flag string    locally-registered flag
error: unknown flag: --global-flag
```


The same thing happens when we pass --help:


```console
$ testcmd --help
Usage:
  testcmd [flags]

Flags:
      --global-flag string   globally-registered flag
  -h, --help                 help for testcmd
      --local-flag string    locally-registered flag
```


This is because Cobra also uses the global flags when generating usage and help text. This can be circumvented by doing-it-yourself, again.


```go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const use = "testcmd"

var (
	globalFlagTarget string
	localFlagTarget  string
)

func init() {
	pflag.StringVar(&globalFlagTarget, "global-flag", globalFlagTarget, "globally-registered flag")
}

func NewLocalFlagSet() *pflag.FlagSet {
	fs := pflag.NewFlagSet(use, pflag.ContinueOnError)
	fs.StringVar(&localFlagTarget, "local-flag", localFlagTarget, "locally-registered flag")
	return fs
}

func NewTestCmd() *cobra.Command {
	localFlagSet := NewLocalFlagSet()
	cmd := &cobra.Command{
		Use:                use,
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			// parse our local flag set
			if err := localFlagSet.Parse(args); err != nil {
				cmd.Usage()
				fatal(err)
			}
			// --help
			help, err := localFlagSet.GetBool("help")
			if err != nil {
				fatal(fmt.Errorf(`"help" flag is non-bool, programmer error, please correct`))
			}
			if help {
				cmd.Help()
				return
			}
			// print the flag values
			fmt.Printf("globalFlagTarget: %q\n", globalFlagTarget)
			fmt.Printf("localFlagTarget: %q\n", localFlagTarget)
		},
	}
	localFlagSet.BoolP("help", "h", false, fmt.Sprintf("help for %s", cmd.Name()))

	// ugly, but necessary, because Cobra's default UsageFunc and HelpFunc pollute the flagset with global flags
	const usageFmt = "Usage:\n  %s\n\nFlags:\n%s"
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine(), localFlagSet.FlagUsagesWrapped(2))
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), usageFmt, cmd.UseLine(), localFlagSet.FlagUsagesWrapped(2))
	})

	return cmd
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func main() {
	cmd := NewTestCmd()
	if err := cmd.Execute(); err != nil {
		fatal(err)
	}
}
```


And now things are more as we expect.


```console
$ testcmd --global-flag hello
Usage:
  testcmd [flags]

Flags:
  -h, --help                help for testcmd
      --local-flag string   locally-registered flag
error: unknown flag: --global-flag

$ testcmd --help
Usage:
  testcmd [flags]

Flags:
  -h, --help                help for testcmd
      --local-flag string   locally-registered flag
```


**Alternative solutions to piecewise-DIY-overrides are highly welcomed** (if someone wants to write a Cobra replacement that meets our needs while managing state more cleanly, I won't stop you). As all core components will need to solve this one way or another, a centralized utility library for working with Cobra would be useful, at the very least.


### Use a flags struct

The migration from flags to versioned config files will be much easier if you first centralize where your target flag values, registrations, and deprecations happen. If you have a single structure definition that contains all of the component's flag-targeted values, you can focus on moving fields from this structure into your versioned configuration API.

The Kubelet uses a structure called KubeletFlags, with an associated `func (f *KubeletFlags) AddFlags(fs *pflag.FlagSet)` to handle flag registrations and deprecations. Note that AddFlags does not register global flags from third-party libraries; it is only concerned with flags in the KubeletFlags structure.

We also recommend that defaulting behavior for flags be clearly separated from the flag registrations. Kubelet initializes flag defaults when constructing a new KubeletFlags, and re-uses these values when registering flags. This makes it easy to see which defaults are applied, which makes it easy to migrate those defaults to versioned config. This also prevents AddFlags from overriding values in the flags struct, in the event that you need to modify values before registering flags. 

Finally, the Kubelet offers a function for validating the flags structure. You may choose to centralize validation here, which will make it easier to migrate that validation to your versioned config, and will also elevate configuration errors to sooner in the component's lifecycle. **It is important to point out that burying flag validation in application logic is an anti-pattern that should be avoided whenever possible. Given the opportunity, components should be refactored to centralize validation.** The Kubelet unfortunately falls into the "validation all over the place" trap, and will eventually need to be refactored to centralize validation.

KubeletFlags currently contains some flags that are only registered on specific operating systems (e.g. Windows). These fields are prefixed with the name of the OS (e.g. `Windows*`) and registrations are handled by OS-specific implementations of the addOSFlags method (managed via Go build tags).

This excerpt from the Kubelet's flag code gives the general structure:


<table>
  <tr>
   <td>cmd/kubelet/app/options/options.go
   </td>
  </tr>
  <tr>
   <td>type KubeletFlags struct {
<p>
	KubeletConfigFile string
<p>
	… 
<p>
}
<p>
func NewKubeletFlags() *KubeletFlags {
<p>
	return &KubeletFlags{
<p>
		// apply defaults here
<p>
	}
<p>
}
<p>
func (f *KubeletFlags) AddFlags(fs *pflag.FlagSet) {
<p>
	f.addOSFlags(fs)
<p>
	fs.StringVar(&f.KubeletConfigFile, "config", f.KubeletConfigFile, "…")
<p>
	… 
<p>
}
<p>
func ValidateKubeletFlags(f *KubeletFlags) error {
<p>
	// validate here, return error if validation fails, nil otherwise
<p>
	… 
<p>
}
   </td>
  </tr>
  <tr>
   <td>cmd/kubelet/app/options/osflags_windows.go
   </td>
  </tr>
  <tr>
   <td>// +build windows
<p>
… 
<p>
func (f *KubeletFlags) addOSFlags(fs *pflag.FlagSet) {
<p>
	// add windows flags here
<p>
}
   </td>
  </tr>
  <tr>
   <td>cmd/kubelet/app/options/osflags_other.go
   </td>
  </tr>
  <tr>
   <td>// +build !windows
<p>
…
<p>
func (f *KubeletFlags) addOSFlags(fs *pflag.FlagSet) {
<p>
	// noop
<p>
}
   </td>
  </tr>
</table>



## Create component's config API group

Components should expose versioned Kubernetes-style configuration APIs. This section explains how to do so. 

As explained in the [API Conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md) doc, Kubernetes API objects consist of a canonical internal version, and multiple external versions. In a given release, it is possible to convert between any of the external versions by first converting to the internal version, then converting from the internal version to the target version. 

All versions live together in the same API group. Typically, an API group's source tree consists of a top-level directory that implements the package for the internal version, and subdirectories corresponding to the packages that implement each external version. There are usually a few additional files that implement various utilities. Finally, there will be generated files (omitted from the below file tree) for conversions, deep-copies, and defaulter registration. 

For example, the file hierarchy of the Kubelet's _kubelet.config.k8s.io_ API group looks like this:


```
- pkg/kubelet/apis/kubeletconfig
| - fuzzer // utility for fuzzing kubelet.config.k8s.io objects
| | - fuzzer.go
|
| - scheme // utility for scheme and codecs (serializations and conversions)
| | - scheme.go
| | - scheme_test.go // round trip tests that use the fuzzer
|
| - v1beta1 // implementation of v1beta1 external type
| | - defaults.go // v1beta1 defaults
| | - doc.go // various build tags that trigger code generation
| | - register.go // functions for registering API with a scheme
| | - types.go // v1beta1 versions of kubelet.config.k8s.io objects
|
| - validation // utility for validating kubelet.config.k8s.io objects
| | - validation.go
|
| - doc.go // various build tags that trigger code generation
| - helpers.go // utility functions
| - helpers_test.go // tests for utility functions
| - register.go // 
| - types.go
```


**When creating your component's API group, please refer to upstream Kubernetes for the most up-to-date example of how it should look, and assign [@mtaufen](https://github.com/mtaufen) or [@liggitt](https://github.com/liggitt) to review your PR.** A few important points to remember are:



*   The API group should be named _{component}.config.k8s.io_, where _{component}_ is the name of your component. It is conventional to also name the directory containing the API group _{component}config_, e.g. _kubeletconfig_.
*   You should start by creating only a _v1alpha1_ external config version, and migrating fields from your flags structure to the objects in this version. Loading config from a file should be considered an alpha feature until you are confident enough in your _v1alpha1_ version to move it to beta. 
*   If your configuration can contain relative file paths, these paths should be resolved relative to the location to the config file when loaded. The Kubelet has a utility function (KubeletConfigurationPathRefs in file _helpers.go_) that enumerates these fields.
*   Once you have created your API group (and whenever you update your config structs), you can run make clean\_generated; make generated\_files to produce the generated conversions, deep-copies, and defaulter registrations.
*   At the beginning, most componentconfig APIs will only load a single object, though this object may be a composition of subobjects. E.g. the Kubelet just loads a KubeletConfiguration, which is a composition of subobjects defined in the _types.go_ files. 
*   Isomorphic objects must exist in both the internal and external versions for conversions to be generated.
*   Only the internal, canonical type needs to be validated. External-versioned config files must be converted to this representation and then validated before your component uses the configuration. 
*   You should not treat nil and empty container types (maps and slices) as semantically different in your API, as this has caused issues with serializers in the past. In practice, this means all map and slice fields should be +optional and omitempty (see [Optional vs. Required](https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#optional-vs-required)). In general, it is rare to have any required config fields, as config versions should ship with functional defaults.
*   You don't have to worry about supporting loading config from Proto yet, but you should not rule this out as an eventual possibility.
*   Try to keep the order of the fields the same between the internal and external types, the defaulter, and the flag registrations. This makes the code much easier to read and maintain.
*   Make component configs modular and composable. Share structs for the same functionality in different components (think of PodSpec shared in many of the core resources).
*   **TODO:** please add more if necessary


## Make it possible to parse flags into the internal config object

In order to maintain command-line compatibility, it must still be possible to parse flags for a period of time after they become available in your config API. When a field moves from your flags structure to your config structure, you should update the corresponding flag registration to target the internal config structure. It helps to keep all flag registrations close to each other. 

For example, Kubelet provides some additional utilities in its options package for targeting KubeletConfiguration values via flags: a constructor (similar to NewKubeletFlags) that returns a default KubeletConfiguration and a function for registering flags that target this config.


<table>
  <tr>
   <td>cmd/kubelet/app/options/options.go
   </td>
  </tr>
  <tr>
   <td>func NewKubeletConfiguration() (*kubeletconfig.KubeletConfiguration, error) {
<p>
	scheme, _, err := kubeletscheme.NewSchemeAndCodecs()
<p>
	if err != nil {
<p>
		return nil, err
<p>
	}
<p>
	versioned := &v1beta1.KubeletConfiguration{}
<p>
	scheme.Default(versioned)
<p>
	config := &kubeletconfig.KubeletConfiguration{}
<p>
	if err := scheme.Convert(versioned, config, nil); err != nil {
<p>
		return nil, err
<p>
	}
<p>
	applyLegacyDefaults(config)
<p>
	return config, nil
<p>
}
<p>
func AddKubeletConfigFlags(fs *pflag.FlagSet, c *kubeletconfig.KubeletConfiguration) {
<p>
	// register flags here, in the same style as in KubeletFlags.AddFlags
<p>
	… 
<p>
}
   </td>
  </tr>
</table>



## Improve/fix defaults between flags and versioned config

Note in the previous example from the Kubelet, there is a call to a function called applyLegacyDefaults. With the move to versioned config files, each API version can have its own set of default values. Flags implicitly constitute their own "version," so it is possible to have different defaults when you load config from a file versus when you load config from flags. NewKubeletConfiguration is specifically constructing a config object to be targeted by flags, so it modifies the values to contain the defaults associated with flags. 

The ability to separate defaults across versions allowed us to use better defaults in _v1beta1_ (see [Secure Kubelet's componentconfig defaults while maintaining CLI compatibility](https://github.com/kubernetes/kubernetes/pull/59666)) than the Kubelet's flag API, without breaking compatibility. If you have defaults in your component's flag API that you would like to change, this is the opportunity to do so.


## Incrementally migrate flags to your config API

For many flags, you will simply be able to:



1. Cut the field from your flags struct.
2. Paste the field into both the internal and versioned config structs.
3. Ensure the field tags are correct on the versioned config struct (json, omitempty, etc.).
4. Move the default into the versioned defaulter, if the old default should be kept.  \
If the versioned default should be different than the old default value, add the new default to the versioned defaulter and the old default into applyLegacyDefaults.
5. Move the flag registration to the AddFlags function for your internal config struct. 

Some, however, will need more work:



*   If a flag embeds structure in a string format (e.g. a list or map), you should use the appropriate language structures (e.g. Go slice or map) to represent that structure in your config API, which will facilitate writing JSON and YAML files. For backwards compatibility, you can write a shim to parse the flag into the structured field. See these PRs for examples:
    *   [Make feature gates loadable from a map[string]bool](https://github.com/kubernetes/kubernetes/pull/53025)
    *   [Lift embedded structure out of ManifestURLHeader field](https://github.com/kubernetes/kubernetes/pull/54643)
    *   [Lift embedded structure out of eviction-related KubeletConfiguration fields](https://github.com/kubernetes/kubernetes/pull/54823)
    *   [ColonSeparatedMultimapStringString: allow multiple Set invocations with default override](https://github.com/kubernetes/kubernetes/pull/55254)
*   If a flag enables or configures alpha or experimental features which do not have an associated [feature gate](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/util/feature/feature_gate.go) (see _[kube\_features.go](https://github.com/kubernetes/kubernetes/blob/master/pkg/features/kube_features.go)_ for a list of gates), you must add a gate or graduate the feature to beta before moving the flag to versioned config. The [API changes policy](https://github.com/kubernetes/community/blob/master/contributors/devel/api_changes.md) allows for alpha fields in beta or GA-versioned config, but only if the behavior configured via the fields is guarded by a feature gate and disabled by default (see [Adding Unstable Features to Stable Versions](https://github.com/kubernetes/community/blob/master/contributors/devel/api_changes.md#adding-unstable-features-to-stable-versions)). Note that if you add an alpha field to beta or GA-versioned config, you permanently exhaust the name you choose for the field; if you change the name before the field graduates from alpha, you must tombstone the old name and never reuse it.
*   If a field requires a [non-zero](https://golang.org/ref/spec#The_zero_value) default value, but the zero-value is _still_ a valid option, it should carry a pointer type in the versioned config struct. This allows the defaulter to differentiate between omitted fields and fields explicitly configured to the zero value. For many fields, especially paths, it is sufficient to simply use the zero value as the default.
    *   Try to avoid pointer fields on the internal type, so that you need fewer nil checks in your code. Let the defaulter handle nils on the external type. You may need to add a [conversion function](https://github.com/kubernetes/kubernetes/blob/e8388e035b1f08d00dd2bf811d9587c6b7589ace/staging/src/k8s.io/apimachinery/pkg/apis/meta/v1/conversion.go) to facilitate pointer to non-pointer conversion generation.
*   If a field is a nilable container type (slice or map) and requires a non-empty default, users must always explicitly set the field to achieve non-default values. If specifying an empty container type was previously valid (e.g. to disable), you will need to provide an explicit substitute for doing so. Since [some serializers (e.g. proto)](https://github.com/kubernetes/kubernetes/issues/43203) do not differentiate between nil and empty container types, neither can we.
    *   See [Add 'none' option to EnforceNodeAllocatable](https://github.com/kubernetes/kubernetes/pull/59515) as an example.

Finally, some flags should not yet be migrated to versioned config. These are unsolved problems, and you should wait to migrate until they can be solved with a consistent pattern across all core components (for more detail, see _Remaining work_).



*   Some flags are only registered in builds for specific operating systems (e.g. the Kubelet's --windows-service flag). We have not settled on how these sorts of fields should be represented in versioned config files, so please refrain from migrating these flags for now.
*   Some flags specify instance-specific values (e.g. the Kubelet's --hostname-override flag). It is impossible to share the same instance-specific value between multiple instances of a component (e.g. you wouldn't want two nodes with the same hostname). If there is (or you suspect) a use case where you want to share the same config source between multiple instances of a component (e.g. if you want to deliver config files via a ConfigMap), you should refrain from migrating these flags for now.

**Note: Do not use unversioned types in your versioned config; stick to language primitives and types from versioned APIs. That said, even using types from other versioned APIs carries risk: you'll have to respond if that API version is deprecated.**


## Component Bootstrap

This section describes the general steps for bootstrapping a component to the point that it knows it has a valid internal-versioned configuration to run with. This should be sufficient for most components. This section begins with the Run method on the Cobra command, and ends with validation of a fully-resolved internal config object.


### Initial Flags Parse

The first thing your component should do is parse its command line into an instance of the flags struct and the internal-verisoned config struct, with defaults applied prior to parsing flags. This is shown in the example in the _Take back control of the command-line API_ section. If you have validation for the flags struct, this is an appropriate place to perform it.


### Load config file, convert to internal version, resolve relative paths

Once the initial flags parse is complete, check whether the user provided a path to --config. If so, load the file at that path (if relative, --config should be resolved relative to the Kubelet's working directory) and unmarshal it via the API machinery's UniversalDecoder, which should perform defaulting and conversion to the internal type. For example:


<table>
  <tr>
   <td>pkg/kubelet/kubeletconfig/util/codec/codec.go
   </td>
  </tr>
  <tr>
   <td>// DecodeKubeletConfiguration decodes a serialized KubeletConfiguration to the internal type
<p>
func DecodeKubeletConfiguration(kubeletCodecs *serializer.CodecFactory, data []byte) (*kubeletconfig.KubeletConfiguration, error) {
<p>
	// the UniversalDecoder runs defaulting and returns the internal type by default
<p>
	obj, gvk, err := kubeletCodecs.UniversalDecoder().Decode(data, nil, nil)
<p>
	if err != nil {
<p>
		return nil, fmt.Errorf("failed to decode, error: %v", err)
<p>
	}
<p>
	internalKC, ok := obj.(*kubeletconfig.KubeletConfiguration)
<p>
	if !ok {
<p>
		return nil, fmt.Errorf("failed to cast object to KubeletConfiguration, unexpected type: %v", gvk)
<p>
	}
<p>
	return internalKC, nil
<p>
}
   </td>
  </tr>
</table>


Next, any fields in the config that specify file paths should be resolved relative to the location of the config file. The Kubelet has a helper that returns pointers to the path fields for a given config (and also a test to detect when new fields are added).


<table>
  <tr>
   <td>pkg/kubelet/apis/kubeletconfig/helpers.go
   </td>
  </tr>
  <tr>
   <td>// KubeletConfigurationPathRefs returns pointers to all of the KubeletConfiguration fields that contain filepaths.
<p>
// You might use this, for example, to resolve all relative paths against some common root before
<p>
// passing the configuration to the application. This method must be kept up to date as new fields are added.
<p>
func KubeletConfigurationPathRefs(kc *KubeletConfiguration) []*string {
<p>
	paths := []*string{}
<p>
	paths = append(paths, &kc.StaticPodPath)
<p>
	paths = append(paths, &kc.Authentication.X509.ClientCAFile)
<p>
	paths = append(paths, &kc.TLSCertFile)
<p>
	paths = append(paths, &kc.TLSPrivateKeyFile)
<p>
	paths = append(paths, &kc.ResolverConfig)
<p>
	return paths
<p>
}
   </td>
  </tr>
</table>


These pointers can be used to resolve relative paths when loading the config file:


<table>
  <tr>
   <td>pkg/kubelet/kubeletconfig/configfiles/configfiles.go
   </td>
  </tr>
  <tr>
   <td>// resolveRelativePaths makes relative paths absolute by resolving them against `root`
<p>
func resolveRelativePaths(paths []*string, root string) {
<p>
	for _, path := range paths {
<p>
		// leave empty paths alone, "no path" is a valid input
<p>
		// do not attempt to resolve paths that are already absolute
<p>
		if len(*path) > 0 && !filepath.IsAbs(*path) {
<p>
			*path = filepath.Join(root, *path)
<p>
		}
<p>
	}
<p>
}
   </td>
  </tr>
</table>



### Enforce flag precedence

If you were able to move all of your command-line flags to versioned config (e.g. you had no flags in the OS-specific, instance-specific, or alpha-not-feature-gated categories), you may not need this step. If you have flags remaining, however, you will need to incrementally migrate them into your API without breaking backwards compatibility. 

Since moving a field to the config structs implicitly adds a default value for that field, you must override this value with the corresponding flag value as long as the flag is set on the command line. Otherwise, simply upgrading a component to a version that migrates a flag to versioned config could break the command-line API. See [#56171](https://github.com/kubernetes/kubernetes/issues/56171) or [Graduating KubeletFlags subfields to KubeletConfiguration](https://docs.google.com/document/d/18-MsChpTkrMGCSqAQN9QGgWuuFoK90SznBbwVkfZryo) for more detail.

The Kubelet achieves this by constructing a flag set that can parse the entire command line, but only populates the config struct as a result of parsing. All non-config flags registrations target mock values, while config flags target real values.

The mock values for global flags are generated by the below helper in the Kubelet, which substitutes values with noop Set operations. NoOp is implemented by [k8s.io/apiserver/pkg/util/flag/noop.go](https://github.com/kubernetes/apiserver/blob/master/pkg/util/flag/noop.go).


<table>
  <tr>
   <td>cmd/kubelet/app/server.go
   </td>
  </tr>
  <tr>
   <td>// newFakeFlagSet constructs a pflag.FlagSet with the same flags as fs, but where
<p>
// all values have noop Set implementations
<p>
func newFakeFlagSet(fs *pflag.FlagSet) *pflag.FlagSet {
<p>
	ret := pflag.NewFlagSet("", pflag.ExitOnError)
<p>
	ret.SetNormalizeFunc(fs.GetNormalizeFunc())
<p>
	fs.VisitAll(func(f *pflag.Flag) {
<p>
		ret.VarP(flag.NoOp{}, f.Name, f.Shorthand, f.Usage)
<p>
	})
<p>
	return ret
<p>
}
   </td>
  </tr>
</table>


We mock component flags by simply targeting a throwaway flags struct. The config flags are simply registered via the AddKubeletConfigFlags function.

Though it is usually sufficient to enforce precedence at flag granularity, for some map fields you may instead wish to enforce precedence at key-value pair granularity. It is generally simpler to avoid doing so, but the FeatureGates field is one example where the Kubelet chooses a piecewise precedence semantic: the key-value pairs from the command line and config file are merged, with the command-line pairs taking precedence. This was done specifically to enable feature rollout via the alpha Dynamic Kubelet Config feature, even when some feature gates are set on the command-line.


<table>
  <tr>
   <td>cmd/kubelet/app/server.go
   </td>
  </tr>
  <tr>
   <td>// Remember original feature gates, so we can merge with flag gates later
<p>
original := kc.FeatureGates
<p>
// re-parse flags
<p>
if err := fs.Parse(args); err != nil {
<p>
	return err
<p>
}
<p>
// Add back feature gates that were set in the original kc, but not in flags
<p>
for k, v := range original {
<p>
	if _, ok := kc.FeatureGates[k]; !ok {
<p>
		kc.FeatureGates[k] = v
<p>
	}
<p>
}
   </td>
  </tr>
</table>



### Validate the config

At this point, you should have a configuration that is ready for validation:



*   A config file was loaded, if specified.
*   Relative paths in the config file were resolved relative to the location of the config file.
    *   Note that relative paths from flags are implicitly relative to the Kubelet's working directory, as this is the default behavior of Go's file system utilities.
*   Flag precedence has been enforced, including feature gate merging.

The next step is to validate this configuration with the validation functions that accompany your API group. If your validation incorporates feature gates, remember to set the feature gates from the config before validating. You can use the [FeatureGate.SetFromMap](https://github.com/kubernetes/kubernetes/blob/52ed0368f8d076236ada19b09828f2f9e2ebb6ef/staging/src/k8s.io/apiserver/pkg/util/feature/feature_gate.go#L84) function to set the gates directly from the field in your config object.

**🎉🎉Congrats, your component now has a versioned config file API! 🎉🎉**


# Remaining work

This section enumerates known issues and shares my opinion on the way forward. The below is all open to discussion.


## Versioned config for third-party flag values

One item we have not discussed in detail is how to migrate flags from third-party libraries to your versioned config file API. 

It is, of course, best if you don't expose these flags at all, but in some cases (e.g. _glog_), doing so is infeasible. In these cases, you can provide a field in your config API for each third-party value and manually plumb the field's value through a flag set with the corresponding flag registered to it.

**Be careful! If a third-party library removes a flag in an update (say, their deprecation policy doesn't align with ours), you'll still be on the hook for maintaining the behavior of fields in your config API.**


## Unsolved problems


### Per-instance configuration

We left instance-specific parameters as flags, because a single config source may need to be shared across multiple instances. We should not, however, abandon these parameters to the command-line. 

The most obvious solution, to me, seems to be the introduction of an --instance-config flag, which accepts a file that contains the instance-specific parameters in a format defined by the component's config API group. This way, instance-specific parameters can be covered by the versioned config file API.


### OS-specific configuration

We left operating-system-specific parameters as flags, because we haven't really had a discussion on how best to represent these in our config APIs. Similar to per-instance configuration, we should not abandon these parameters to the command-line.

I think that a simple and flexible solution is to prefix OS-specific fields with the name of the OS:



*   This helps avoid the complexity of having a top-level substructure for each supported OS, or OS specific substructures of general substructures.
*   These fields should always be optional and omitempty, so that they can be omitted in environments where they are not necessary. 
*   If we see common fields, it should be relatively easy to add a non-prefixed field that works across multiple supported operating systems.
*   We'll likely need to ensure that defaulting and validation only process OS-specific fields for the OS we are currently running on.
*   It would be really nice if we could have API machinery that could mark which fields are supported on which OSes.


# TODOs

**TODO: Dig through the following and identify things I missed in this doc:**

Original dynamic config proposal discussion, esp. points about API policies: https://github.com/kubernetes/kubernetes/pull/29459#issuecomment-271990251 

Major related PRs:



*   [Graduate kubeletconfig API group to beta](https://github.com/kubernetes/kubernetes/pull/53833)
*   [Deprecate KubeletConfiguration flags](https://github.com/kubernetes/kubernetes/pull/60148)

Other Related PRs (massaging KubeletConfiguration object to beta quality, fixing flag precedence, testing, correct loading behavior, etc):



*   [Make feature gates loadable from a map[string]bool](https://github.com/kubernetes/kubernetes/pull/53025)
*   [Mulligan: Remove deprecated and experimental fields from KubeletConfiguration](https://github.com/kubernetes/kubernetes/pull/53088)
*   [Move --enable-cusom-metrics to KubeletFlags and mark it deprecated](https://github.com/kubernetes/kubernetes/pull/54154)
*   [Move runtime-related flags from KubeletConfiguration to KubeletFlags](https://github.com/kubernetes/kubernetes/pull/54160)
*   [Lift embedded structure out of ManifestURLHeader field](https://github.com/kubernetes/kubernetes/pull/54643)
*   [Lift embedded structure out of eviction-related KubeletConfiguration fields](https://github.com/kubernetes/kubernetes/pull/54823)
*   [ColonSeparatedMultimapStringString: allow multiple Set invocations with default override](https://github.com/kubernetes/kubernetes/pull/55254)
*   [Move 'alpha' KubeletConfiguration fields that aren't feature-gated and self-registration fields to KubeletFlags](https://github.com/kubernetes/kubernetes/pull/55562)
*   [Kubelet: Relative paths in local config file](https://github.com/kubernetes/kubernetes/pull/55648)
*   [Add kubeletconfig round trip test](https://github.com/kubernetes/kubernetes/pull/55961)
*   [seccomp is an alpha feature and not feature gated](https://github.com/kubernetes/kubernetes/pull/55983)
*   [Kubelet flags take precedence over config from files/ConfigMaps](https://github.com/kubernetes/kubernetes/pull/56097)
*   [flag precedence redo](https://github.com/kubernetes/kubernetes/pull/56995)
*   [Refactor kubelet config controller bootstrap process](https://github.com/kubernetes/kubernetes/pull/57488)
*   [Explicit kubelet flags](https://github.com/kubernetes/kubernetes/pull/57613)
*   [Fix PodCIDR flag: defaults come from the object, not as literal args to the flag function](https://github.com/kubernetes/kubernetes/pull/57621)
*   [Replace --init-config-dir with --config](https://github.com/kubernetes/kubernetes/pull/57624)
*   [e2e node framework can generate a base kubelet config file](https://github.com/kubernetes/kubernetes/pull/57638)
*   [More default fixups for Kubelet flags](https://github.com/kubernetes/kubernetes/pull/57770)
*   [Move some old security controls to KubeletFlags and mark them deprecated](https://github.com/kubernetes/kubernetes/pull/57851)
*   [Turn KubeletConfigFile on in e2enode tests](https://github.com/kubernetes/test-infra/pull/5637)
*   [Removal of KubeletConfigFile feature gate: Step 1](https://github.com/kubernetes/kubernetes/pull/58760)
*   [Removal of KubeletConfigFile feature gate: Step 2 (stop setting KubeletConfigFile)](https://github.com/kubernetes/test-infra/pull/6490)
*   [Removal of KubeletConfigFile feature gate: Step 3 (final)](https://github.com/kubernetes/kubernetes/pull/58978)
*   [--generate-kubelet-config-file=true is now default](https://github.com/kubernetes/test-infra/pull/6654)
*   [Fix PodPidsLimit and ConfigTrialDuration on internal KubeletConfig type](https://github.com/kubernetes/kubernetes/pull/59062)
*   [Add 'none' option to EnforceNodeAllocatable](https://github.com/kubernetes/kubernetes/pull/59515)
*   [remove CAdvisorPort from KubeletConfiguration struct](https://github.com/kubernetes/kubernetes/pull/59580)
*   [Bury KubeletConfiguration.ConfigTrialDuration for now](https://github.com/kubernetes/kubernetes/pull/59628)
*   [Secure Kubelet's componentconfig defaults while maintaining CLI compatibility](https://github.com/kubernetes/kubernetes/pull/59666)
*   [Ignore 0% and 100% eviction thresholds](https://github.com/kubernetes/kubernetes/pull/59681)
*   [expunge the word 'manifest' from Kubelet's config API](https://github.com/kubernetes/kubernetes/pull/60314)
