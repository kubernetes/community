# Persist cluster-wide data in ConfigMap

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
