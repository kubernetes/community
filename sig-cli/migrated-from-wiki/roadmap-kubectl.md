# kubectl roadmap

`kubectl` is the Kubernetes CLI.

If you'd like to contribute, please read the [conventions](/contributors/devel/sig-cli/kubectl-conventions.md) and familiarize yourself with [existing commands](http://kubernetes.io/docs/user-guide/kubectl-overview/).

**Owner:** @kubernetes/kubectl

**Label:** [component/kubectl](https://github.com/kubernetes/kubernetes/labels/component%2Fkubectl)

**Motivation:** [kubectl brainstorm](https://docs.google.com/document/d/1tPrTL5Fi8BjlIK_XwNz-W260ll2ZYflrdbtnwE1PQoE/edit?pli=1#)

### Add new commands / subcommands / flags
* [Simplify support for multiple files](https://github.com/kubernetes/kubernetes/issues/24649)
  * Manifest that can specify multiple files / http(s) URLs
  * [Default manifest](https://github.com/kubernetes/kubernetes/issues/3268) (ala Dockerfile or Makefile)
  * Unpack archive (tgz, zip) and then invoke “-f” on that directory
  * URL shortening via default URL prefix
* [Imperative `set` commands](https://github.com/kubernetes/kubernetes/issues/21648)
* [`view` commands](https://github.com/kubernetes/kubernetes/issues/29679)
* [Support `run --edit` and `create --edit`](https://github.com/kubernetes/kubernetes/issues/18064)
* [More `kubectl create <sub-command>`](https://github.com/kubernetes/kubernetes/issues/25382)
* [Support `--dry-run` for every mutation](https://github.com/kubernetes/kubernetes/issues/11488)
* kubectl commands aliases
  * [Allow user defined aliases for resources and commands](https://github.com/kubernetes/kubernetes/issues/18023)
  * [Suggest possibly matching kubectl commands](https://github.com/kubernetes/kubernetes/issues/25180)
* Improve `kubectl run`
  * Make generated objects more discoverable: suggest the user to do `kubectl get all` to see what's generated ([extend `all` to more resources](https://github.com/kubernetes/kubernetes/issues/22337))
  * [Make it optional to specify name (auto generate name from image)](https://github.com/kubernetes/kubernetes/issues/2643)
  * [Make `kubectl run --restart=Never` creates Pods (instead of Jobs)](https://github.com/kubernetes/kubernetes/issues/24533)
* Create commands/flags for common get + template patterns (e.g. getting service IP address)
* [Implement `kubectl cp`](https://github.com/kubernetes/kubernetes/issues/13776) to copy files between containers and local for debugging
* `kubectl rollout`
  * [Add `kubectl rollout start` to show how to start a rollout](https://github.com/kubernetes/kubernetes/issues/25142)
  * [Add `kubectl rollout status`](https://github.com/kubernetes/kubernetes/issues/25235)
* Scripting support
  * [wait](https://github.com/kubernetes/kubernetes/issues/1899)
  * [watch / IFTTT](https://github.com/kubernetes/kubernetes/issues/5164)
* [Add `kubectl top`](https://github.com/kubernetes/kubernetes/issues/11382) which lists resource metrics.

### Alternative interfaces

* Create a terminal based console, ref [docker console](https://github.com/dustinlacewell/console) ([video](https://www.youtube.com/watch?v=wSzZxbDYgtY))
* [Add `kubectl sh`, an interactive shell](https://github.com/kubernetes/kubernetes/issues/25385), or make a kubectlshell in contrib and make bash completion part of it (ref [pythonshell](https://gist.github.com/bprashanth/9a3c8dfbba443698ddd960b8087107bf))
* Think about how/whether to invoke generation commands such as `kubectl run` or `kubectl create configmap` in bulk, declaratively, such as part of the `apply` flow.
* [ChatOps](https://www.pagerduty.com/blog/what-is-chatops/) bot -- such as [kubebot](https://github.com/harbur/kubebot) (add to tools documentation)

### Improve help / error messages / output
* Make kubectl functionality more discoverable
  * [Overhaul kubectl help](https://github.com/kubernetes/kubernetes/issues/16089)
    * ~~[Print "Usage" at the bottom](https://github.com/kubernetes/kubernetes/issues/7496)~~
    * Add keywords (critical words) to help
    * List valid resources for each command
    * Make short description of each command more concrete; use the same language for each command
    * Link to docs ([kubernetes.io/docs](http://kubernetes.io/docs))
    * [Update `kubectl help` descriptions and examples from docs](https://github.com/kubernetes/kubernetes/issues/25290)
    * Embed formatting and post-process for different media (terminal, man, github, etc.)
    * [Suppress/hide global flags](https://github.com/kubernetes/kubernetes/issues/23402)
    * ~~[Categorize kubectl commands or list them in alphabetical order]~~(https://github.com/kubernetes/kubernetes/issues/21585)
    * [Implement search in `kubectl help`](https://github.com/kubernetes/kubernetes/issues/25234)
  * [Suggest next/alternative commands](https://github.com/kubernetes/kubernetes/issues/19736)
  * [Add a verbosity flag that explains all the things that it's doing](https://github.com/kubernetes/kubernetes/issues/25272)
  * ~~[Fix incomplete kubectl bash completion](https://github.com/kubernetes/kubernetes/issues/25287)~~
* Improve error messages (note that not all of these problems are in kubectl itself)
  * [when kubectl doesn’t know what cluster to talk to](https://github.com/kubernetes/kubernetes/issues/24420)
  * ~~[non-existent namespace produces obscure error](https://github.com/kubernetes/kubernetes/issues/15542)~~
  * [line numbers with validation errors](https://github.com/kubernetes/kubernetes/issues/12231)
  * [invalid lines with validation errors](https://github.com/kubernetes/kubernetes/issues/6132)
  * [malformed inputs produce misleading error messages](https://github.com/kubernetes/kubernetes/issues/9012)
  * [non-yaml/json produces obscure error](https://github.com/kubernetes/kubernetes/issues/8838)
  * [error messages for non-existent groups/types](https://github.com/kubernetes/kubernetes/issues/19530)
    * Suggest resource type when not provided (e.g. `kubectl get my-pod-name` should suggest running `kubectl get pod/my-pod-name`)
    * [errors for other non-existent resources](https://github.com/kubernetes/kubernetes/issues/6703)
  * [lack of apiVersion/kind produces confusing error messages](https://github.com/kubernetes/kubernetes/issues/6439)
  * [update validation errors could be more informative](https://github.com/kubernetes/kubernetes/issues/8668)
  * [field validation errors could be more helpful](https://github.com/kubernetes/kubernetes/issues/10534)
  * [field errors should use json field names](https://github.com/kubernetes/kubernetes/issues/3084)
  * [clearer error for bad image/registry](https://github.com/kubernetes/kubernetes/issues/7960)
  * [no error for illegal scale](https://github.com/kubernetes/kubernetes/issues/11148)
  * [deletion timeout doesn't provide any details](https://github.com/kubernetes/kubernetes/issues/19427)
  * [service creation timeout doesn't provide any details](https://github.com/kubernetes/kubernetes/issues/4860)
  * [create secret with invalid data has obscure error message](https://github.com/kubernetes/kubernetes/issues/10309)
  * [--all-namespaces error is unclear](https://github.com/kubernetes/kubernetes/issues/15834)
  * [exec has unclear errors](https://github.com/kubernetes/kubernetes/issues/9944)
  * [logs has misleading errors](https://github.com/kubernetes/kubernetes/issues/6376)
  * [improve error reporting by adding URLs](https://github.com/kubernetes/kubernetes/issues/5551)
  * Improve jsonpath / gotemplate error messages (it's tricky to get the path just right)
  * [error message for user with no permissions is extremely cryptic](https://github.com/kubernetes/kubernetes/issues/26909)
* [Cleanup `kubectl get/describe` output](https://github.com/kubernetes/kubernetes/issues/20941)
  * [Clarify kubectl get/describe service output](https://github.com/kubernetes/kubernetes/issues/22702)
* [Define and document command conventions for users](https://github.com/kubernetes/kubernetes/issues/25388)

### Bug fix
* Fix [apply](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aopen+label%3Acomponent%2Fkubectl+label%3Akind%2Fbug+apply), [edit](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aopen+label%3Acomponent%2Fkubectl+label%3Akind%2Fbug+edit), and [validate](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aopen+label%3Acomponent%2Fkubectl+label%3Akind%2Fbug+validate) bugs

### Installation / Release
* `gcloud` should enable kubectl bash completion when installing `kubectl`
* [Pipe-to-sh to install kubectl](https://github.com/kubernetes/kubernetes/issues/25386)
* [Static build of kubectl for containers](https://github.com/kubernetes/kubernetes/issues/23708) ([we have it](https://git.k8s.io/kubernetes/examples/kubectl-container), but it's not part of the release)

### Others
* [Move functionality to server](https://github.com/kubernetes/kubernetes/issues/12143)
* [Eliminate round-trip conversion of API objects in kubectl](https://github.com/kubernetes/kubernetes/issues/3955)
* [Move preferences out of kubeconfig](https://github.com/kubernetes/kubernetes/issues/10693)
* And then add more preferences
  * Enable/disable explanatory mode (see [kploy output](https://github.com/kubernauts/kploy))
  * Permanently disable warnings once displayed
  * Default labels as columns
  * Default `--record`, `--save-config`, etc.
* [Overhaul cluster-related commands](https://github.com/kubernetes/kubernetes/issues/20605)
  * [Delete cluster from `kubectl config`](https://github.com/kubernetes/kubernetes/issues/25601)
* ~~["kubectl-only Ubernetes": enabe kubectl to target any one of many clusters](https://github.com/kubernetes/kubernetes/issues/23492)~~
