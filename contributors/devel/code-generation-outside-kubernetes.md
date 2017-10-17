# Code generation (from the experience of developing service-catalog)

## Code generator binaries
There are several libraries in the Kubernetes repo that are responsible for
various code generation operations. Those go based libraries are currently
located at cmd/libs/go2idl/. In each directory, there is a main.go file
that currently has hard coded an internal path for boilerplate.go.txt.
However, there now exists a [repo-infra](https://github.com/kubernetes/repo-infra)
repository that can be used instead. These files serve mostly as wrappers for
executing and utilizing gengo.

## Code generation arguments
In Kubernetes, the code generation libraries exist in the cmd/libs/go2idl
directory. Most of the code generation execution handling is in the
Makefile.generated_files and is in the .generated_files target. The rule
complexity is due to Kubernetes generating a DAG in order to determine
which files need to be rebuilt and code regenerated for.

Each code generation binary supports different arguments. The arguments of
importance are:
--extra-peer-dirs - defines additional locations to provide current operations
--bounding-dirs - list of import paths for binding the types
--input-dirs - list of import paths to get input types from

In service-catalog, the makefile has been significantly simplified and one
can see the parameters sent to each code generator here:

https://github.com/kubernetes-incubator/service-catalog/blob/59cd78081f6cb23aa102df493f5b2471e8e0cdd2/Makefile#L102

Please note the conversion-gen arguments, which include Kubernetes directories.
The reason for this is because the Kubernetes types have custom decoders defined
and are inherited on any project specific types. These codecs being necessary is
also why update-codecgen.sh must be run as well:

https://github.com/kubernetes-incubator/service-catalog/blob/59cd78081f6cb23aa102df493f5b2471e8e0cdd2/Makefile#L116

## Marking files for code generation
Special flags are required to tag each package eligible for code generation.
Each starts with a "+k8s:" prefix. The convention is to add these tags into
a doc.go file in each version variant. For example, you can see the tags
listed in one such file here:

https://github.com/kubernetes-incubator/service-catalog/blob/59cd78081f6cb23aa102df493f5b2471e8e0cdd2/pkg/apis/servicecatalog/v1alpha1/doc.go

There are various options that are passed along with the flag that changes
generation behavior. These options are described in either main.go for a
given generator or within gengo itself.

## Potential pitfall
When importing Kubernetes api it's important to remember to import the
correct version of the api. Always align the api with the same version
package that is being built:

https://github.com/kubernetes-incubator/service-catalog/blob/59cd78081f6cb23aa102df493f5b2471e8e0cdd2/pkg/apis/servicecatalog/types.go#L20
https://github.com/kubernetes-incubator/service-catalog/blob/59cd78081f6cb23aa102df493f5b2471e8e0cdd2/pkg/apis/servicecatalog/v1alpha1/types.go#L20
