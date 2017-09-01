# Service IP Configuration

## Abstract

A proposal for making the configuration of services that are exposed to the outside world more flexible. This will be achieved by making settings to the underlying cloud provider available in the annotations of the Kubernetes service.

## Motivation

Kubernetes currently create a random IP and associate it with the service, this doesn't allow for creating more advanced setups for applications that require it.

The IP is managed by the cloud using a default configuration that might not match the needs of the user and require manual configuration of the underlying cloud infrastructure in order to have the desired access to the service.

There are several settings a user may want to be able to control:
 * Having the IP static.
 * Giving an internal DNS name in addition to the numeric IP.
 * Sharing the same IP between multiple services.

This document describes the interface to configure the underlying cloud IP allocation from a Kubernetes service level.

## Use Cases

 * Configuring a static IP that will not be changed in case of a restart or downtime of the service.
   It is possible that an IP will be freed by the cloud provider if the Kubernetes cluster is shutdown for maintanance (e.g. changing virtual machine size).
 * Configuraing a DNS name for the service to have a consistent and memorable name to access the service.
   The user have a DNS server that is configured to redirect to an internal CNAME to avoid having to reconfigure in case the IP changes (see case #1).
 * Allowing multiple services (and pods) to share the same IP but a different port.
   If there are several pods that provide different but related functionality on different pods they can be configured to use the same IP and be routed by different services.

## Design

A Kubernetes service that expose a port to the outside world is using the cloud provider's API to create a public IP address and route the traffic from this address to the load balancer that in turn route it to the pods that match the service's selector.
The request to the cloud provider's API is sent with the default configuration and cannot be changed through Kubernetes.

By extending the service definition Kubernetes can forward the settings down to the cloud provider's API.

This information is only relevant to the underlying infrastructure and doesn't affect the Kubernetes objects and their behavior. For that porpuse we suggest to use annotations.

The following extensions are suggested as new annotations to service:

 1. `service.alpha.kubernetes.io/dns-label-name` - configure a label to be used in the last part of the DNS name given to the IP address.  
    e.g. myservice.westus.cloudapp.azure.com
 2. `service.alpha.kubernetes.io/static-ip` - a boolean flag to request whether the allocated IP will be static.

When a new service is created, the module responsible for creating the IP address allocation will read the annotations and if they are supported by the cloud provider it will pass the relevant parameters to the cloud provider's API.

## Cloud Provider Support

### AWS

Amazon AWS supports setting a static IP address, called elastic IP address.

`https://ec2.amazonaws.com/?Action=AllocateAddress&AUTHPARAMS`

Documentation: [AllocateAddress](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AllocateAddress.html)

### Azure

Microsoft Azure supports setting a static IP and an internal dns label name:

```
PUT /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Network/publicIPAddresses/{publicIPAddressName}?api-version={api-version}
"properties": {
                "publicIPAddressVersion": "IPv4",
                "publicIPAllocationMethod": "Static",
                "dnsSettings": {
                    "domainNameLabel": "traffic-receiver-test"
                }
}
```

Documentation: [Create or update a public IP address](https://docs.microsoft.com/en-us/rest/api/network/virtualnetwork/create-or-update-a-public-ip-address)

