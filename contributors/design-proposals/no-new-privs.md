#Support "no new privileges" in Kubernetes

##Description

In Linux, the `execve` system call can grant more privileges to a newly-created process than its parent process. Considering security issues, since Linux kernel v3.5, there is a new flag named `no_new_privs` added to prevent those new privileges from being granted to the processes.

`no_new_privs` is inherited across `fork`, `clone` and `execve` and can not be unset. With `no_new_privs` set, `execve` promises not to grant the privilege to do anything that could not have been done without the `execve` call.

For more details about `no_new_privs`, please check the Linux kernel document [here](https://www.kernel.org/doc/Documentation/prctl/no_new_privs.txt).

Docker started to support `no_new_privs` option since 1.11. Here is the [link](https://github.com/docker/docker/issues/20329) of the ticket in Docker community to support `no_new_privs` option.

We want to support the creation of containers with `no_new_privs` enabled in Kubernetes, which will make the Kubernetes cluster more safe. Here is the [link](https://github.com/kubernetes/kubernetes/issues/38417) of the ticket in Kubernetes community to track this proposal.


##Current implementation

###Support in Docker

Since Docker 1.11, user can specify `--security-opt` to enable `no_new_privs` while creating containers, e.g. `docker run --security-opt=no-new-privileges busybox`

For program client, Docker provides an object named `ContainerCreateConfig` defined in package `github.com/docker/engine-api/types` to config container creation parameters. In this object, there is a string array `HostConfig.SecurityOpt` to specify the security options. Client can utilize this field to specify the arguments for security options while creating new containers.

###Support in OCI runtimes

Since version 0.3.0 of the OCI runtime specification, a user can specify the `noNewPrivs` boolean flag in the configuration file.

More details of OCI implementation can be checked [here](https://github.com/opencontainers/runtime-spec/pull/290).

###SecurityContext in Kubernetes

Kubernetes defines `SecurityContext` for `Container` and `PodSecurityContext` for `PodSpec`. `SecurityContext` objects define the related security options for Kubernetes containers, e.g. selinux options.

While creating a container, kubelet parses the security context object and formats the security option strings for Docker. The security options strings will finally be inserted into `ContainerCreateConfig.HostConfig.SecurityOpt` and passed to Docker. Different Kubernetes runtimes now are using different methods to parse and format the security option strings:
* method `#getSecurityOpts` in `docker_mager_xxxx.go` for Docker runtime
* method `#getContainerSecurityOpts` in `docker_container.go` for CRI


##Proposal to support "no new privileges"

To support "no new privileges" options in Kubernetes, it is proposed to make the following changes:

###Changes of SecurityContext objects

Add a new bool type field named `noNewPrivileges` to both `SecurityContext` definition and `PodSecurityContext` definition:
* `noNewPrivileges=true` in `PodSecurityContext` means that all the containers in the pod should be run with `no-new-privileges` enabled. This should be a pod level control of `no-new-privileges` flag.
* `noNewPrivileges` in `SecurityContext` is a container level control of `no-new-privileges` flag, and can override the pod level `noNewPrivileges` setting.

By default,  `noNewPrivileges` is `false`.

The change of security context API objects requires the update of corresponding Kubernetes documents, need to submit another PR to track this.

###Changes of docker runtime

When parsing the new `SecurityContext` object, kubelet has to take care of `noNewPrivileges` field from security context objects. Once `noNewPrivileges` is `true`, kubelet needs to change `#getSecurityOpts` method in `docker_manager_xxx.go` to add `no-new-privileges` option to `ContainerCreateConfig.HostConfig.SecurityOpt`

###Changes of CRI runtime

When parsing the new `SecurityContext` object, kubelet has to take care of `noNewPrivileges` field from security context objects. Once `noNewPrivileges` is `true`, kubelet needs to change `#getContainerSecurityOpts` method in `docker_container.go` to add `no-new-privileges` option to `ContainerCreateConfig.HostConfig.SecurityOpt`

###Changes of kubectl

This is an additional proposal for kubectl. To improve kubectl user experience, we can add a new flag for kubectl command named `--security-opt`. This flag allows user to create pod with security options configured when using `kubectl run` command. For example, if user issues command like `kubectl run busybox --image=busybox --security-opt=no-new-privileges -- top`, kubernetes shall create a pod with `noNewPrivileges` enabled.

If the proposal of kubectl changes is accepted, the patch can also be submitted as a separate PR.
