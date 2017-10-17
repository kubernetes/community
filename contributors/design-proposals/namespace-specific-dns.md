#Description

Multi-tenant is a software architecture which describes that one single instance of software serves multiple customers at the same time. This architecture provides isolated view of the whole system to each of the customer so that they co-exist in the same system and won’t affect, or see, each other. 

Kubernetes as a container orchestrator platform, uses namespace to realize multi-tenant architecture. But for now, there is only one cluster level DNS server deployed in the kubernetes environment, which means that pods can find/resolve-to the service that is from another namespace.

To enhance the multi-tenant support, this proposal suggests to support namespace specific DNS servers. The use case for this proposal is a kubernetes operator who is providing a multi-tenant deployment environment based on namespaces. Each namespace belongs to one tenant. The operator can deploy separate DNS servers for different namespaces, so each tenant gets its own namespace and DNS server. In this case, users from other namespaces won't find/resolve-to the services defined in our own namespace. This will help keep the privacy of the multi-tenant environment.


#Current implementation of KubeDNS

At present, kubernetes deploys KubeDNS as a service. This service contains 3 containers(since kubernetes v1.4)
* dnsmasq container: provides basic DNS query services based on DNS record stored in kubedns container
* kubedns container: the core of KubeDNS, does the following things:
  * start a SkyDNS server
  * monitor the service changes in k8s cluster and store kubernetes service DNS record in memory
* exechealthz container: provides health check for the above two containers

The kubedns container monitors the service changes in kubernetes cluster and updates DNS records. Generally, it provides the following methods:
* setServicesStore: this method is invoked when launching a new DNS server. The method creates a pair of `cache.Store` and `cache.Controller` objects to monitor the changes of the kubernetes services from all the namespaces in the cluster and registers three callback functions to take corresponding actions for managing the services’ DNS records:
  * newService: callback function, when there is a new service created in the cluster, this method will be called and KubeDNS will create a new DNS record for the new service information, like service name, namespace info and service IP.
  * updateService: callback function, when an existing service is updated in the cluster, this method will be called and KubeDNS will update the DNS record of this service.
  * removeService: callback function, when an existing service is removed in the cluster, this method will be called and KubeDNS will clear the DNS record of this service.
* setEndpointsStore: this method is invoked when launching a new DNS server. The method creates a pair of `cache.Store` and `cache.Controller` objects to monitor the changes of the kubernetes endpoints from all the namespaces in the cluster and registers one callback function to take corresponding actions for adding or updating the endpoints’ DNS records:
* handleEndpointAdd: callback function, when there is a new endpoint created or an existing endpoint updated in the cluster, this method will be called and KubeDNS will create relevant DNS records for the endpoint information.
Start: this is the public method of kubeDNS to lunch DNS servers. In this method, the services `cache.Controller` and endpoints `cache.Controller` created in `#setServicesStore` and `#setEndpointsStore` will be run as go routines to start monitor the resources changes from all the namespaces.

#Proposal to support namespace specific DNS server

To support namespace specific DNS server in kubernetes, there are at least two parts of changes shall be made: KubeDNS API changes and Namespace specific DNS client config.

##KubeDNS API changes.

* A new KubeDNS flag named `--dns-namespace` is provided to KubeDNS

Kubernetes runs `kube-dns` executable as the major container process of kubedns container. We can add a new flag named `--dns-namespace` to specify the namespace that this single DNS server serves when running `kube-dns`. In addition, this DNS server will also serve the request for namespace `default`. In other words, the namespace specified DNS server will only store DNS records and provide DNS query service for the kubernetes services/endpoints belong to `default` namespace or the namespace defined by `--dns-namespace`.

Once `--dns-namespace` flag is provided to `kube-dns`, instead of monitoring the changes from the whole cluster’s kubernetes services/endpoints, this single DNS server will only monitor the changes of kubernetes services/endpoints that belong to the in-scope namespaces —— dns-namespace & default-namespace. Therefore, we can modify the implementation of `#setServicesStore` and `#setEndpointsStore` to register the callback methods only for the services/endpoints from specific namespaces. As a result, instead of just creating one pair of `cache.Store` and `cache.Controller` objects to list & watch resource changes from all the namespaces in `#setServicesStore` and `#setEndpointsStore` respectively, we would create two pairs of the store and controller objects under this situation. 1 pair of the store and controller objects lists & watches the resource changes from `default` namespace, while the other pair monitors the resource changes from the namespace defined by `--dns-namespace`.

Accordingly, the `#Start` method now will need to start 4 go routines for the controllers. 2 controllers to watch service changes from default-namespace and dns-namespace, while 2 other controllers for endpoint changes.

As a result, if the above changes are made to KubeDNS and launch multi-DNS server with the `--dns-namespace` flag, each single DNS server will only provide DNS service for the kubernetes services from restricted/scoped namespaces.

There is already some POC code of this proposal available [here](https://github.com/duglin/kubernetes/commit/15fbe749cb8df7a637e1365059ac987b9c89a089).

* `--dns-namespace` flag is not provided to KubeDNS

Under this situation, the behavior of KubeDNS should remain the same as before, which means, `#setServicesStore` and `#setEndpointsStore` should only create the controllers to monitor the resource changes from all the namespaces.

* Namespace specific DNS client config

Once we enable a namespace specific DNS server in a kubernetes cluster, we shall make kubelet aware of multiple DNS servers. When creating pods from different namespaces, kubelet can add different DNS client config to let different pods point to different DNS servers.

To implement this, kubernetes should allow namespace rather than kubelet to supply the added DNS client config. We already submitted a proposal for adding this feature in K8S community, please refer to this [link](https://github.com/kubernetes/features/issues/75) to check the details of the design, and also the patch of the changes is available [here](https://github.com/ardnaxelarak/kubernetes/commit/53ec8dbae38b15e4236996dc7857b19648c64b7b).

