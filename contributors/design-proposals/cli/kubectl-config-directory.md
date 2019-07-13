# Kubectl Configuration Directory

Status: Pending

Version: Alpha

Implementation Owner: TBD

## Motivation

Managing configuration for multiple clusters, users and accounts can be complex. Some users will use separate config files, providing them as needed with the [`KUBECONFIG`](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/#the-kubeconfig-environment-variable) environment variable. Others will put multiple configurations in a single file, and use the `--context` option. Users who prefer to manage a single configuration file may quickly find it becomes very large and complex to manage.

This design document proposes a feature which will be very familiar to Linux/Unix users, which allows many small configuration files to be loaded from a single, well-defined location.

## Proposal

The proposal is that we provide a 'configuration directory', at the following location:

```
$HOME/.kube/config.d
```

On startup, `kubectl` will load the standard configuration file, and then merge into its in-memory structure each file which exists in the `config.d` directory.

This is a fairly idiomatic approach for Linux/Unix software (`apt`, `nginx`, `apache`, `systemd` etc). It would be a non-breaking feature, and would support the use cases listed below.

## User Experience

### Use Cases

The following use-cases have been considered:

#### Rapid Access to Newly Created Clusters

Consider the following project:

[terraform-aws-openshift](https://github.com/dwmkerr/terraform-aws-openshift)

This project creates a Kubernetes Cluster. One the cluster has been created, it is convenient to offer the user a way to access it. At the moment, the only options are:

1. Generate a config file, ask the user to download it, and ask them to refer to it with the `KUBECONFIG` variable. This has the drawback of forcing the user to decide where to store the file, and remember that location whenever they access the cluster.
2. Ask the user to modify their config file. This has the drawback of forcing the user to modify a complex, and sometimes very large file.
3. Programmatically update the config file. This is the approach I have used the most frequently. However, it can be complex as well. Below is an example of programmatic manipulation of the config file:

```
#!/usr/bin/env bash
cluster_name="mycluster"
context_name="mycontext"
user_name="kubernetes-admin"

# Create the location to store certs and keys.
# Define paths to hold credentials.
keys_path="$HOME/.kube/${cluster_name}-credentials"
mkdir -p ${keys_path}
config_path="${keys_path}/config"
certificate_authority_path="${keys_path}/${cluster_name}.ca.crt"
client_certificate_path="${keys_path}/${user_name}.crt"
client_key_path="${keys_path}/${user_name}.key"

cluster_dns="$(terraform output api_dns)"
echo "Cluster load balancer is running on: $cluster_dns..."

# Get the kubeconfig from the bastion.
scp root@$(terraform output bastion_public_dns):~/.kube/config $config_path

# Rip out the data we need.
cat $config_path \
    | sed -n -E 's/.*certificate-authority-data:[ +](.*)/\1/p' \
    | base64 --decode > $certificate_authority_path
cat $config_path \
    | sed -n -E 's/.*client-certificate-data:[ +](.*)/\1/p' \
    | base64 --decode > $client_certificate_path
cat $config_path \
    | sed -n -E 's/.*client-key-data:[ +](.*)/\1/p' \
    | base64 --decode > $client_key_path

# Create the cluster config.
kubectl config set-cluster ${cluster_name} \
    --server=https://${cluster_dns}:6443 \
    --certificate-authority=${certificate_authority_path} \

# Create the user config.
kubectl config set-credentials ${user_name} \
    --client-certificate=${client_certificate_path} \
    --client-key=${client_key_path}

# Create a context pointing the user to the cluster.
kubectl config set-context ${context_name} \
    --cluster=${cluster_name} \
    --namespace=default \
    --user=${user_name}
```

In this example, data is pulled from Terraform to find a bastion address. From the user's host, certificates and keys are copied from the bastion. Finally, we use the `kubectl config set-*` commands to programmatically update the configuration file.

Whilst this approach works, it is somewhat complex, and _still_ leads the problem of having very large configuration files, which can be hard to manage.

With the approach given above, the following would happen:

```sh
scp root@$(terraform output bastion_public_dns):~/.kube/config.d/mycluster $config_path
```

And a generated configuration file would be added to the user's `config.d` directory. If the user wants to remove this configuration at any stage, it becomes trivial to simply delete the appropriate config file, rather than manipulating the large 'master' config file. The user can also choose to structure their configuration as they prefer, as the number of clusters they manage grows.

## Implementation

Essentially we would start with the code in:

[client-go/blob/master/tools/clientcmd/loader.go](https://github.com/kubernetes/client-go/blob/master/tools/clientcmd/loader.go)

And simply add the logic to loop over a given folder, recursively, searching for config files. For each file found, we augment the current in-memory structure.

### Client/Server Backwards/Forwards compatibility

I do not believe this will have compatibility issues, *except* in the case where users are *already* storing configuration in the `config.d` directory for their own convenience. In this case, the users will have the config loaded by default, which would not have been the case before. I believe this can be mitigated with an 'opt in' flag, or simply by well documenting the feature and letting it move through the alpha/beta lifecycle changes, giving users time to become aware of it.

## Alternatives considered

I have not considered other alternatives as at this stage there _already_ are alternatives for managing configuration (such as separate files managed with `KUBECONFIG` or simply one large config file).
