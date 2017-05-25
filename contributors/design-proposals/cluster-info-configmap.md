# Persist cluster-wide data in ConfigMap

### Terminology

The following terminology will be used throughout the rest of this design proposal. The meanings are excerpted from the sig-service-catalog [terminology page](https://github.com/kubernetes-incubator/service-catalog/blob/master/terminology.md).

**Service**: Running code that is made available for use by an application. Traditionally, services are available via HTTP REST endpoints, but this is not a requirement.

**Service Broker**: An endpoint that manages a set of services. Responsible for translating Service Catalog activities (like provision, bind, unbind, deprovision) into appropriate actions for the service.

**Service Catalog**: An endpoint that manages a set of registered Service Brokers and (2) the list of services that are available for instantiation from those Service Brokers. For more information please see [sig-service-catalog](https://github.com/kubernetes-incubator/service-catalog).


## Abstract

We want to allow cluster-wide data to be persisted for general purposes, currently for us (service catalog) is to focus on the ability to identify different clusters that access our service broker with a 'cluster-id'.

## Motivation

Currently there is no way for service catalog api server to identify the different clusters that access our service brokers. The goal of this design is to have a unify storage for cluster related information. The information should be retained and survive between component failures. 

## Detailed design

A configmap `cluster-info` is created in the `default` namespace at the starup of the controller. The initial content of the map will be as the following:

```
apiVersion: v1
data:
  cluster-id: uuid-of-the-cluster
kind: ConfigMap
metadata:
  name: cluster-info
  namespace: default
```

The controller attempts to read the configmap at startup, if the configmap does not already exist, the controller generates a cluster-id and creates a new configmap. To avoid discrepancy, the controller also reconciles with the configmap periodically.



