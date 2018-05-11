## Workflow

![Git workflow](git_workflow.png)

### 1 Fork in the cloud

1. Visit https://github.com/kubernetes/kubernetes
2. Click `Fork` button (top right) to establish a cloud-based fork.

### 2 Clone fork to local storage

Per Go's [workspace instructions][go-workspace], place Kubernetes' code on your
`GOPATH` using the following cloning procedure.

[go-workspace]: https://golang.org/doc/code.html#Workspaces

Define a local working directory:

```sh
# If your GOPATH has multiple paths, pick
# just one and use it instead of $GOPATH here.
# You must follow exactly this pattern,
# neither `$GOPATH/src/github.com/${your github profile name/`
# nor any other pattern will work.
working_dir=$GOPATH/src/k8s.io
```

> If you already do Go development on github, the `k8s.io` directory
> will be a sibling to your existing `github.com` directory.

Set `user` to match your github profile name:

```sh
user={your github profile name}
```

Both `$working_dir` and `$user` are mentioned in the figure above.

Create your clone:

```sh
mkdir -p $working_dir
cd $working_dir
git clone https://github.com/$user/kubernetes.git
# or: git clone git@github.com:$user/kubernetes.git

cd $working_dir/kubernetes
git remote add upstream https://github.com/kubernetes/kubernetes.git
# or: git remote add upstream git@github.com:kubernetes/kubernetes.git

# Never push to upstream master
git remote set-url --push upstream no_push

# Confirm that your remotes make sense:
git remote -v
```

### 3 Branch

Get your local master up to date:

```sh
cd $working_dir/kubernetes
git fetch upstream
git checkout master
git rebase upstream/master
```

Branch from it:
```sh
git checkout -b myfeature
```

Then edit code on the `myfeature` branch.

#### Build
The following section is a quick start on how to build Kubernetes locally, for more detailed information you can see [kubernetes/build](https://git.k8s.io/kubernetes/build/README.md).
The best way to validate your current setup is to build a small part of Kubernetes. This way you can address issues without waiting for the full build to complete. To build a specific part of Kubernetes use the `WHAT` environment variable to let the build scripts know you want to build only a certain package/executable.

```sh
make WHAT=cmd/{$package_you_want}
```

*Note:* This applies to all top level folders under kubernetes/cmd.

So for the cli, you can run:

```sh
make WHAT=cmd/kubectl
```

If everything checks out you will have an executable in the `_output/bin` directory to play around with.

*Note:* If you are using `CDPATH`, you must either start it with a leading colon, or unset the variable. The make rules and scripts to build require the current directory to come first on the CD search path in order to properly navigate between directories.

```sh
cd $working_dir/kubernetes
make
```

To remove the limit on the number of errors the Go compiler reports (default
limit is 10 errors):
```sh
make GOGCFLAGS="-e"
```

To build with optimizations disabled (enables use of source debug tools):

```sh
make GOGCFLAGS="-N -l"
```

To build binaries for all platforms:

```sh
make cross
```

#### Install etcd

```sh
cd $working_dir/kubernetes

# Installs in ./third_party/etcd
hack/install-etcd.sh

# Add to PATH
echo export PATH="\$PATH:$working_dir/kubernetes/third_party/etcd" >> ~/.profile
```

#### Test

```sh
cd $working_dir/kubernetes

# Run all the presubmission verification. Then, run a specific update script (hack/update-*.sh)
# for each failed verification. For example:
#   hack/update-gofmt.sh (to make sure all files are correctly formatted, usually needed when you add new files)
#   hack/update-bazel.sh (to update bazel build related files, usually needed when you add or remove imports)
make verify

# Alternatively, run all update scripts to avoid fixing verification failures one by one.
make update

# Run every unit test
make test

# Run package tests verbosely
make test WHAT=./pkg/api/helper GOFLAGS=-v

# Run integration tests, requires etcd
# For more info, visit https://git.k8s.io/community/contributors/devel/testing.md#integration-tests
make test-integration

# Run e2e tests by building test binaries, turn up a test cluster, run all tests, and tear the cluster down
# Equivalent to: go run hack/e2e.go -- -v --build --up --test --down
# Note: running all e2e tests takes a LONG time! To run specific e2e tests, visit:
# https://git.k8s.io/community/contributors/devel/e2e-tests.md#building-kubernetes-and-running-the-tests
make test-e2e
```

See the [testing guide](/contributors/devel/testing.md) and [end-to-end tests](/contributors/devel/e2e-tests.md)
for additional information and scenarios.

Run `make help` for additional information on these make targets.

### 4 Keep your branch in sync

```sh
# While on your myfeature branch
git fetch upstream
git rebase upstream/master
```

Please don't use `git pull` instead of the above `fetch` / `rebase`. `git pull`
does a merge, which leaves merge commits. These make the commit history messy
and violate the principle that commits ought to be individually understandable
and useful (see below). You can also consider changing your `.git/config` file via
`git config branch.autoSetupRebase always` to change the behavior of `git pull`.

### 5 Commit

Commit your changes.

```sh
git commit
```
Likely you go back and edit/build/test some more then `commit --amend`
in a few cycles.

### 6 Push

When ready to review (or just to establish an offsite backup or your work),
push your branch to your fork on `github.com`:

```sh
git push -f ${your_remote_name} myfeature
```

### 7 Create a pull request

1. Visit your fork at https://github.com/$user/kubernetes
2. Click the `Compare & Pull Request` button next to your `myfeature` branch.
3. Check out the pull request [process](/contributors/guide/pull-requests.md) for more details and
   advice.

_If you have upstream write access_, please refrain from using the GitHub UI for
creating PRs, because GitHub will create the PR branch inside the main
repository rather than inside your fork.

#### Get a code review

Once your pull request has been opened it will be assigned to one or more
reviewers.  Those reviewers will do a thorough code review, looking for
correctness, bugs, opportunities for improvement, documentation and comments,
and style.

Commit changes made in response to review comments to the same branch on your
fork.

Very small PRs are easy to review.  Very large PRs are very difficult to review.
At the assigned reviewer's discretion, a PR may be switched to use
[Reviewable](https://reviewable.k8s.io) instead.  Once a PR is switched to
Reviewable, please ONLY send or reply to comments through Reviewable.  Mixing
code review tools can be very confusing.

#### Squash and Merge

Upon merge (by either you or your reviewer), all commits left on the review
branch should represent meaningful milestones or units of work.  Use commits to
add clarity to the development and review process.

Before merging a PR, squash any _fix review feedback_, _typo_, _merged_, and
_rebased_ sorts of commits.

It is not imperative that every commit in a PR compile and pass tests
independently, but it is worth striving for.

In particular, if you happened to have used `git merge` and have merge
commits, please squash those away: they do not meet the above test.

A nifty way to manage the commits in your PR is to do an [interactive
rebase](https://git-scm.com/book/en/v2/Git-Tools-Rewriting-History),
which will let you tell git what to do with every commit:

```sh
git fetch upstream
git rebase -i upstream/master
```

For mass automated fixups (e.g. automated doc formatting), use one or more
commits for the changes to tooling and a final commit to apply the fixup en
masse. This makes reviews easier.
