# Generation and release cycle of clientset

Client-gen is an automatic tool that generates [clientset](/contributors/design-proposals/api-machinery/client-package-structure.md#high-level-client-sets) based on API types. This doc introduces the use of client-gen, and the release cycle of the generated clientsets.

## Using client-gen

The workflow includes three steps:

**1.** Marking API types with tags: in `pkg/apis/${GROUP}/${VERSION}/types.go`, mark the types (e.g., Pods) that you want to generate clients for with the `// +genclient` tag. If the resource associated with the type is not namespace scoped (e.g., PersistentVolume), you need to append the `// +genclient:nonNamespaced` tag as well.

The following `// +genclient` are supported:

- `// +genclient` - generate default client verb functions (*create*, *update*, *delete*, *get*, *list*, *update*, *patch*, *watch* and depending on the existence of `.Status` field in the type the client is generated for also *updateStatus*).
- `// +genclient:nonNamespaced` - all verb functions are generated without namespace.
- `// +genclient:onlyVerbs=create,get` - only listed verb functions will be generated.
- `// +genclient:skipVerbs=watch` - all default client verb functions will be generated **except** *watch* verb.
- `// +genclient:noStatus` - skip generation of *updateStatus* verb even thought the `.Status` field exists.

In some cases you want to generate non-standard verbs (eg. for sub-resources). To do that you can use the following generator tag:

- `// +genclient:method=Scale,verb=update,subresource=scale,input=k8s.io/api/extensions/v1beta1.Scale,result=k8s.io/api/extensions/v1beta1.Scale` - in this case a new function `Scale(string, *v1beta.Scale) *v1beta.Scale` will be added to the default client and the body of the function will be based on the *update* verb. The optional *subresource* argument will make the generated client function use subresource `scale`. Using the optional *input* and *result* arguments you can override the default type with a custom type. If the import path is not given, the generator will assume the type exists in the same package.

In addition, the following optional tags influence the client generation:

- `// +groupName=policy.authorization.k8s.io` – used in the fake client as the full group name (defaults to the package name),
- `// +groupGoName=AuthorizationPolicy` – a CamelCase Golang identifier to de-conflict groups with non-unique prefixes like `policy.authorization.k8s.io` and `policy.k8s.io`. These would lead to two `Policy()` methods in the clientset otherwise (defaults to the upper-case first segement of the group name).

**2a.** If you are developing in the k8s.io/kubernetes repository, you just need to run hack/update-codegen.sh.

**2b.** If you are running client-gen outside of k8s.io/kubernetes, you need to use the command line argument `--input` to specify the groups and versions of the APIs you want to generate clients for, client-gen will then look into `pkg/apis/${GROUP}/${VERSION}/types.go` and generate clients for the types you have marked with the `genclient` tags. For example, to generated a clientset named "my_release" including clients for api/v1 objects and extensions/v1beta1 objects, you need to run:

``` 
$ client-gen --input="api/v1,extensions/v1beta1" --clientset-name="my_release"
```

**3.** ***Adding expansion methods***: client-gen only generates the common methods, such as CRUD. You can manually add additional methods through the expansion interface. For example, this [file](https://git.k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/core/internalversion/pod_expansion.go) adds additional methods to Pod's client. As a convention, we put the expansion interface and its methods in file ${TYPE}_expansion.go. In most cases, you don't want to remove existing expansion files. So to make life easier, instead of creating a new clientset from scratch, ***you can copy and rename an existing clientset (so that all the expansion files are copied)***, and then run client-gen.

## Output of client-gen

- clientset: the clientset will be generated at `pkg/client/clientset_generated/` by default, and you can change the path via the `--clientset-path` command line argument.

- Individual typed clients and client for group: They will be generated at `pkg/client/clientset_generated/${clientset_name}/typed/generated/${GROUP}/${VERSION}/`

## Released clientsets

If you are contributing code to k8s.io/kubernetes, try to use the generated clientset [here](https://git.k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset).

If you need a stable Go client to build your own project, please refer to the [client-go repository](https://github.com/kubernetes/client-go).

We are migrating k8s.io/kubernetes to use client-go as well, see issue [#35159](https://github.com/kubernetes/kubernetes/issues/35159).
