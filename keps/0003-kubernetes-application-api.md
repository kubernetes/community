---
kep-number: draft-20180115
title: Kubernetes Application API
authors:
  - "@kow3ns"
owning-sig: sig-apps
participating-sigs:
  - sig-apps
  - sig-cli
reviewers:
  - "@kow3ns"
  - "@danielromlein"
  - "@lavalamp"
approvers:
  - "@kow3ns"
  - "@mattfarina"
  - "@prydonius"
  - "@pwittrock"
editor:
creation-date: 2018-01-15
last-updated: 2018-01-15
status: draft
see-also:
replaces:
superseded-by:
---

# Kubernetes Application API

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
 * [Goals](#goals)
 * [Non-Goals](#non-goals)
* [Proposal](#proposal)
 * [User Stories](#user-stories)
    * [Application Creation](#application-creation)
    * [Application View](#application-view)
    * [Application Health](#application-health)
    * [Application Update](#application-update)
    * [Application List](#application-list)
    * [Application Watch](#application-watch)
    * [Application Deletion](#application-deletion)
 * [Implementation](#implementation)
 * [API](#api)
    * [Validation](#validation)
    * [Generation Incrementation](#generation-incrementation)
    * [Namespaces](#namespaces)
    * [Composition (Is A)](#composition-is-a)
    * [Aggregation (Has A)](#aggregation-has-a)
    * [Selecting Components](#selecting-components)
    * [Retrieving Applications](#retrieving-applications)
    * [Retrieving Dependencies](#retrieving-dependencies)
    * [Health Checks](#health-checks)
 * [Application Controller](#application-controller)
    * [Garbage Collection](#garbage-collection)
    * [Installation Status Reporting](#installation-status-reporting)
    * [Health Check Pods](#health-check-pods)
    * [Health Status Reporting.](#health-status-reporting)
    * [Generation Observation](#generation-observation)
    * [Adopting Existing Components](#adopting-existing-components)
 * [kubectl](#kubectl)
 * [UI Considerations](#ui-considerations)
    * [Discovering Application Components and Dependencies](#discovering-application-components-and-dependencies)
    * [Discovering an Object's Application](#discovering-an-objects-application)
    * [Application Status](#application-status)
 * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
 * [Planned](#planned)
* [Drawbacks](#drawbacks)
 * [Increased Memory, CPU, and Network Requirements](#increased-memory-cpu-and-network-requirements)
 * [Increased API Server Storage Requirements](#increased-api-server-storage-requirements)
* [Alternatives](#alternatives)
 * [Use Labeling Conventions](#use-labeling-conventions)
 * [Use Custom Resource Definitions](#use-custom-resource-definitions)
 * [Application Type Aggregation](#application-type-aggregation)
 * [Use the Garbage Collector Instead of an Application Controller](#use-the-garbage-collector-instead-of-an-application-controller)






## Summary
Kubernetes has many primitives for managing workloads (e.g. Pods, ReplicaSets, Deployments, DaemonSets and 
StatefulSets), storage (e.g. PersistentVolumeClaims and PersistentVolumes), and networking (e.g. Services, 
Headless Services, and Ingeresses). When these primitives are aggregated to provide a service to an end user or to 
another system, the whole becomes something more than the individual parts. Instead of a set of loosely coupled 
workloads and their corresponding storage and networking, we have an application. 

Currently, there is no standard way for tools to discover the relationship between the components of an 
application. Many tools use bespoke schemes (generally involving labeling and/or annotations). This has lead to a 
proliferation on non-interoperable tools, a non-uniform experience for application developers and users, and an 
inability of UI (User Interface) designers to surface information about applications to users.

To address these issues, we propose to add an Application Kind as a first class citizen to the apps Group of the 
Kuberentes API. This Kind can be used by tools to communicate that the applications they create are more than just 
a loosely coupled set of API objects.

The use of the Application object is strictly opt-in for tool developers. Tools and users that choose to use the 
Application object in their manifests will benefit from interoperability, a uniform user experience, and first class 
support by Kubernetes UIs.

## Motivation
As an example, consider how [WordPress](https://github.com/WordPress/WordPress), a simple CMS  (Content Management 
System) is created via a manifest. The manifest below creates two StatefulSets (one for a MySQL RDBMs instance and 
another for the WordPress web server), a Headless Services for each StatefulSet, a and load balanced Service for the 
WordPress web server.
 
 ```yaml
 apiVersion: v1
 kind: Service
 metadata:
   name: wordpress-mysql-hsvc
   labels:
     app: wordpress
     component: wordpress-mysql-hsvc
 spec:
   ports:
     - port: 3306
   selector:
     app: wordpress
     component: wordpres-mysql
   clusterIP: None
 ---  
 apiVersion: apps/v1
 kind: StatefulSet
 metadata:
   name: wordpress-mysql
   labels:
     app: wordpress
 spec:
   selector:
     matchLabels:
       app: wordpress
       component: wordpress-mysql
   replicas: 1
   service: wordpress-mysql-hsvc
   template:
     metadata:
       labels:
         app: wordpress
         component: rdbms
     spec:
      
       containers:
       - image: mysql:5.6
         name: mysql
         env:
         - name: MYSQL_ROOT_PASSWORD
           valueFrom:
             secretKeyRef:
               name: mysql-pass
               key: password
         ports:
         - containerPort: 3306
           name: mysql
         volumeMounts:
         - name: mysql-persistent-storage
           mountPath: /var/lib/mysql
    volumeClaimTemplates:
      - metadata:
          name: mysql-persistent-storage
        spec:
          accessModes: [ "ReadWriteOnce" ]
          resources:
            requests:
              storage: 250Gi
 ---
 apiVersion: v1
 kind: Service
 metadata:
   name: wordpress-webserver-svc
   labels:
     app: wordpress
     component: wordpress-webserver-svc
 spec:
   ports:
     - port: 80
   selector:
     app: wordpress
     component: wordpress-webserver
   type: LoadBalancer
 ---
 apiVersion: v1
 kind: Service
 metadata:
   name: wordpress-webserver-hsvc
   labels:
     app: wordpress
     component: wordpress-webserver-hsvc
 spec:
   ports:
     - port: 3306
   selector:
     app: wordpress
     component: wordpress-webserver
   clusterIP: None
 ---
 apiVersion: apps/v1
 kind: StatefulSet
 metadata:
   name: wordpress
   labels:
     app: wordpress
     component: wordpress-webserver
 spec:
   replicas: 1
   service: wordpress-webserver-hsvc
   selector:
     matchLabels:
       app: wordpress
       component: wordpress-webserver
   template:
     metadata:
       labels:
         app: wordpress
         tier: frontend
     spec:
       containers:
       - image: wordpress:4.8-apache
         name: wordpress
         env:
         - name: WORDPRESS_DB_HOST
           value: wordpress-mysql
         - name: WORDPRESS_DB_PASSWORD
           valueFrom:
             secretKeyRef:
               name: mysql-pass
               key: password
         ports:
         - containerPort: 80
           name: wordpress
         volumeMounts:
         - name: wordpress-persistent-storage
           mountPath: /var/www/html
   volumeClaimTemplates:
        - metadata:
            name: mysql-persistent-storage
          spec:
            accessModes: [ "ReadWriteOnce" ]
            resources:
              requests:
                storage: 250Gi
 ```
 
Form an infrastructure management perspective, the manifest above encapsulates everything that is necessary to create 
the workloads and services that compose our WordPress CMS application. However, the fact that these workloads comprise 
an application is lost in translation during creation. Once these manifests are applied, there is no default way 
for users to reconstruct the fact that the whole is more than the sum of its parts.
 
When UIs, either command line or graphical, seek to interact with the application, they can only interact with its 
components, and there is no standard way to indicate that those components comprise an application. As previously 
mentioned, Many tools in the ecosystem attempt to aggregate the components of an application using labels and/or 
annotations, but, as each tool uses a different scheme, an application created by one tool is generally not recognizable 
by another, and UI designers have no standard way to recognize and display an application created by any tool.

Moreover, their is no way to determine if the application is healthy. Even if a UI is able to determine the components 
and dependencies of the application, and even if it is able to determine that the individual components are healthy 
(e.g All Services are created and all Deployments are fully replicated), this does not necessarily imply that the 
application as a whole is able to service requests.


### Goals

1. Provide a standard API for creating, viewing, and managing applications in Kubernetes.
1. Provide a CLI implementation, via kubectl, that interacts with the Application API.
1. Provide installation status and garbage collection for applications.
1. Provide a standard way for applications to surface a basic health check to the UIs.
1. Provide an explicit mechanism for applications to declare dependencies on another application.
1. Promote interoperability among ecosystem tools and UIs by creating a standard that tools MAY implement.

### Non-Goals

1. Create a standard that all tools MUST implement. 
1. Provide a way for UIs to surface metrics from an application.

## Proposal
In order to address the issues discussed in the [previous section](#motivation) we will introduce the Application Kind 
to the apps Group of the Kubernetes API, add the Application Controller to the Controller Manager, and modify 
kubectl to make use the of the Application Kind.

### User Stories

#### Application Creation
As an application maintainer or tool developer, I can create an application that is interoperable with other tools and 
has first class support in Kuberentes UIs.

#### Application View
As an application user or UI designer, I have a standard way to view the components and status of Kubernetes 
applications.

#### Application Health
As an application developer, I can surface a basic health check for UIs to consume, and, as a UI developer, I have a 
standard way to consume application level health checks.

#### Application Update
As an application developer, maintainer, or administrator, I can update the configuration of my application.

#### Application List
As an application developer, maintainer or administrator, I can list the applications that are currently installed.

#### Application Watch
As an application developer, maintainer or administrator, I can watch the applications in a cluster for changes.

#### Application Deletion
As a user, I can delete my application, and when I delete my Application, its corresponding components are deleted.

### Implementation
This section describes the implementation of the feature, including the API, the associated controller, and the CLI 
and UI interaction.

### API
The go code below describes the Application Kind. This code will be added to apps/v1alpha1 Group, and promoted through 
apps/v1beta1 and apps/v1 Groups as stability of the Kind matures.

```go

// ApplicationStatus contains the status of an Application. 
type ApplicationStatus struct {
    // ObservedGeneration is used by the Application Controller to report the last Generation of an Application 
    // that it has observed.
    ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"int64,1,opt,name=observedGeneration"`
    
    // Installed is used by the Application Controller ot report the installed Components and Dependencies of  
    // an Application.
	Installed []string `json:"installed,omitempty" protobuf:"bytes,2,opt,name=installed"`
	
	// Ready is used to report the readiness of an Application. If an Application contains an ApplicationHealthCheck, 
	// this field corresponds directly to the Readiness of the Pod generated from that health check. If the 
	// Application does not contain an ApplicationHealthCheck, this field is set to Unknown.
	Ready ConditionStatus json:`"ready,omitempty" protobuf:"bytes,3,opt,name=ready"`
}

// ApplicationHealthCheck contains a template for a Pod that performs an Application health check. The readiness 
// of the Pod decides the readiness of the Application.
type ApplicationHealthCheck struct {

   // Template is used to specify a Pod whose readiness will be used to determine the readiness of an Application.
   Template v1.PodTemplateSpec
}
// ApplicationSpec is the specification object for an Application. It specifies the Application's components, 
// dependencies, and an optional health check. Its selector can be used to select application components.
type ApplicationSpec struct {
    // Type is the type of the application (e.g. WordPress, MySQL, Cassandra).
    Type string `json:"type" protobuf:"bytes,1,name=type"`

	// Components is a map of the applications components to the kinds created by the application (e.g. Pods, Services, 
	// Deployments, CRDS, etc)
	Components map [string]metav1.GroupVersionKind `json:"components,omitempty" protobuf:"bytes,2,opt,name=components"`
	
	// Dependencies is a list of other Applications that this Application depends on.
	Dependencies []sting `json:"dependencies" protobuf:"bytes,3,opt,name=dependencies"`

	// selector is a label query over kinds that created by the application. It must match the component objects' labels.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector *metav1.LabelSelector `json:"selector" protobuf:"bytes,4,opt,name=selector"`
	
	// HealthCheck is an optional application level health check.
	HealthCheck* ApplicationHealthCheck `json:"healthCheck" protobuf:"bytes,5,opt,name=healthCheck"`

	// Version is an optional version indicator for the application.
	Version string `json:"version,omitempty" protobuf:"bytes,6,opt,name=version"`
	
	// AboutURL is a URL pointing to any additional information about the application.
    AboutURL string `json:"url,omitempty" protobuf:"bytes,7,opt,name=url"`
	
	// Description is a brief string description of the application.
    Description string `json:"description,omitempty" protobuf:"bytes,8,opt,name=description"`
	
	// Maintainers is an optional list of maintainers of the application
	Maintainers []string `json:"maintainers,omitempty" protobuf:"bytes,9,opt,name=maintainers"`
	
	// Keywords is an optional list of key words associated with the application (e.g. MySQL, RDBMS, database).
	Keywords []string `json:"keywords,omitempty" protobuf:"bytes,10,opt,name=keywords"`
	
	// LicenseURL is a an optional URL pointing to any licensing information for the application.
	LicenseURL string `json:"license,omitempty" protobuf:"bytes,11,opt,name=license"`
	
	// DashboardURL is a an optional URL pointing to the applications dashboard.
    DashboardURL string `json:"dashboard,omitempty" protobuf:"bytes,12,opt,name=dashboard"`
	
}

// The Application object acts as an aggregator for components that comprise an application. Its Spec.Components 
// indicate the components that comprise the application and the GroupVersionKinds of the those components. 
// Its SpecSSelector is used to list and watch the Components that correspond to a GVK. All components of an 
// Application should be labeled such the Application's Selector matches. An Applications Spec.Dependencies indicate 
// the other Applications on which the Application depends.
type Application struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	
	// Spec contains the specification for the Application
	Spec ApplicationSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	
	// Status contains the Status of the application with respect to the installed components.
	Status ApplicationStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}
```

The purpose of the Application Kind is to aggregate the components and dependencies of a Kubernetes application such 
that they are explicit and discoverable. An Application's `Spec.Components` is a mapping of the components in an 
application to the GVK (Group Version Kind) used to realize the component. Each component must be labeled such that its 
owning Application's `Spec.Selector` correctly matches the component's labels. Clients of the API Server can use the GVK 
information, along with the label selector, to list or watch the components in an Application. Additionally, each 
component should be annotated with a reference to the Application.

An Application's `Spec.Dependencies` is a list of other Applications on which the application depends. Clients can use 
this field to discover and retrieve other Applications that the containing Application requires to be present in the 
system.

While we attempt to represent many properties of the Application object with explicit fields, we realize that the 
current list is not exhaustive, and we do not endeavor to develop an exhaustive list prior to releasing the 
Application Kind. We expect users to use the fields of ObjectMeta (i.e. Labels and Annotations) for custom use cases. 
When the community forms a consensus that a particular label or annotation has become prevalent across a large 
number of tools, we will promote that metadata to a field. Note that this can be done in a backward compatible manner 
without respect to the stability of the Application Kind.

#### Validation
Upon Application creation, The API Server shall perform the following validation.

1. The API Server shall validate that the Application is a valid ObjectMeta with a name that is a valid DNS subdomain. 
In effect, the API Server will simply validate that the Application is a valid Kubernetes resource.
1. The API Server will validate that the Application's `Spec.Selector` is a valid, non-nil, set-based Selector.
1. If the Application has a `.Spec.HealthCheck`, the API Server will validate that the `Template` of the 
ApplicationHealthCheck is a valid Pod with a readiness check.

The API Server shall perform the following validation upon Application mutation.

1. The API server shall ensure that the new version of the Application is valid using the above definition of a valid 
newly created Application instance. 
1. The API server shall ensure that the Applications `Spec.Selector` has not been mutated. This latter condition has 
been implemented in many places in the Kubernetes API to prevent unintentional orphaning. This is our purpose here.

#### Generation Incrementation

1. Upon creation of an Application object, the API Server shall set the initial value of the `Generation` field to 0.
1. The API Server shall, upon update of an Application object, increment the `Generation` of the Application object by 1. 


#### Namespaces
All components and dependencies of an Application must reside in the same namespace as the Application. Given the 
design of Kubernetes namespaces, we feel this is a reasonable, and desirable, constraint. A user, who has write 
permissions for a particular namespace, can install an Application and all of its components into the same namespace, 
and only those users with the necessary permissions can modify the Application. Additionally, administrators may limit 
the resources consumed by an Application by applying quotas to the namespace in which it resides.

#### Composition (Is A)
The Application Kind is designed to handle the composition of Kubernetes primitive Kinds. This includes CRDs 
(Custom Resource Definitions). The manifest below demonstrates how this can be applied to the WordPress CMS.


```yaml
apiVersion: v1
Kind: Application
metadata:
  name: wordpress
spec:
  type: wordpress
  version: 1.0
  url: 
  selector:
    matchLables:
      app: wordpress
  components:
    - wordpress-mysql-hsvc:
        version: v1
        kind: Service
    - wordpress-mysql:
        group: apps
        version: v1
        kind: StatefulSet
    - wordpress-webserver-hsvc:
        version: v1
        kind: Service
    - wordpress-webserver-svc:
        version: v1
        kind: Service
    - wordpress-mysql:
        group: apps
        version: v1
        kind: StatefulSet
    - wordpress-webserver:
        group: apps
        version: v1
        kind: StatefulSet
---
apiVersion: v1
kind: Service
metadata:
  name: wordpress-mysql-hsvc
  labels:
    app: wordpress
    component: wordpress-mysql-hsvc
  annotations:
    kubernetes.io/application: wordpress
spec:
  ports:
    - port: 3306
  selector:
    app: wordpress
    component: wordpres-mysql
  clusterIP: None
---  
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: wordpress-mysql
  labels:
    app: wordpress
  annotations:
      kubernetes.io/application: wordpress
spec:
  selector:
    matchLabels:
      app: wordpress
      component: wordpress-mysql
  replicas: 1
  service: wordpress-mysql-hsvc
  template:
    metadata:
      labels:
        app: wordpress
        component: rdbms
    spec:
     
      containers:
      - image: mysql:5.6
        name: mysql
        env:
        - name: MYSQL_ROOT_PASSWORD
          valueFrom:
            secretKeyRef:
              name: mysql-pass
              key: password
        ports:
        - containerPort: 3306
          name: mysql
        volumeMounts:
        - name: mysql-persistent-storage
          mountPath: /var/lib/mysql
   volumeClaimTemplates:
     - metadata:
         name: mysql-persistent-storage
       spec:
         accessModes: [ "ReadWriteOnce" ]
         resources:
           requests:
             storage: 250Gi
---
apiVersion: v1
kind: Service
metadata:
  name: wordpress-webserver-svc
  labels:
    app: wordpress
    component: wordpress-webserver-svc
  annotations:
      kubernetes.io/application: wordpress
spec:
  ports:
    - port: 80
  selector:
    app: wordpress
    component: wordpress-webserver
  type: LoadBalancer
---
apiVersion: v1
kind: Service
metadata:
  name: wordpress-webserver-hsvc
  labels:
    app: wordpress
    component: wordpress-webserver-hsvc
  annotations:
      kubernetes.io/application: wordpress
spec:
  ports:
    - port: 3306
  selector:
    app: wordpress
    component: wordpress-webserver
  clusterIP: None
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: wordpress-webserver
  labels:
    app: wordpress
    component: wordpress-webserver
  annotations:
      kubernetes.io/application: wordpress
spec:
  replicas: 1
  service: wordpress-webserver-hsvc
  selector:
    matchLabels:
      app: wordpress
      component: wordpress-webserver
  template:
    metadata:
      labels:
        app: wordpress
        tier: frontend
    spec:
      containers:
      - image: wordpress:4.8-apache
        name: wordpress
        env:
        - name: WORDPRESS_DB_HOST
          value: wordpress-mysql
        - name: WORDPRESS_DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: mysql-pass
              key: password
        ports:
        - containerPort: 80
          name: wordpress
        volumeMounts:
        - name: wordpress-persistent-storage
          mountPath: /var/www/html
  volumeClaimTemplates:
       - metadata:
           name: mysql-persistent-storage
         spec:
           accessModes: [ "ReadWriteOnce" ]
           resources:
             requests:
               storage: 250Gi
```

The components of the WordPress application are made explicit by the `Spec.Components` of the `word-press` Application.
When clients examine the `Spec.Components` of the Application, they have enough information to retrieve and display 
any of the Kubernetes primitives that compose the Application. When any of those components is examined, clients can use 
the name indicated by the `kubernetes.io/application` to retrieve the parent Application object.

#### Aggregation (Has A)
Applications may depend on other applications. Going back to the example of WordPress, deploying a single RDBMs per 
WordPress web server instance is only one way to organize the application. In some instances, it may be desirable to 
create a single, replicated, highly available RDBMS that can used by many instances of the WordPress web server. The 
Application Kind allows users to express an aggregation via its `Spec.Dependencies`.

```yaml
apiVersion: v1
Kind: Application
metadata:
  name: mysql
spec:
  type: mysql
  version: 1.0
  url: 
  selector:
    matchLables:
      app: mysql
  components:
    - mysql-hsvc:
        version: v1
        kind: Service
    - mysql:
        group: apps
        version: v1
        kind: StatefulSet
---
apiVersion: v1
kind: Service
metadata:
  name: mysql-hsvc
  labels:
    app: wordpress
    component: mysql-hsvc
  annotations:
    kubernetes.io/application: mysql
spec:
  ports:
    - port: 3306
  selector:
    app: wordpress
    component: wordpres-mysql
  clusterIP: None
---  
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mysql
  labels:
    app: wordpress
  annotations:
      kubernetes.io/application: mysql
spec:
  selector:
    matchLabels:
      app: wordpress
      component: mysql
  replicas: 1
  service: mysql-hsvc
  template:
    metadata:
      labels:
        app: wordpress
        component: rdbms
    spec:
     
      containers:
      - image: mysql:5.6
        name: mysql
        env:
        - name: MYSQL_ROOT_PASSWORD
          valueFrom:
            secretKeyRef:
              name: mysql-pass
              key: password
        ports:
        - containerPort: 3306
          name: mysql
        volumeMounts:
        - name: mysql-persistent-storage
          mountPath: /var/lib/mysql
   volumeClaimTemplates:
     - metadata:
         name: mysql-persistent-storage
       spec:
         accessModes: [ "ReadWriteOnce" ]
         resources:
           requests:
             storage: 250Gi
--- 
apiVersion: v1
Kind: Application
metadata:
  name: wordpress
spec:
  type: wordpress
  version: 1.0
  url: 
  selector:
    matchLables:
      app: wordpress
  components:
    - wordpress-hsvc:
        version: v1
        kind: Service
    - wordpress-svc:
        version: v1
        kind: Service
    - wordpress-webserver:
        group: apps
        version: v1
        kind: StatefulSet
   dependencies:
    - mysql
---
apiVersion: v1
kind: Service
metadata:
  name: wordpress-svc
  labels:
    app: wordpress
    component: wordpress-svc
  annotations:
      kubernetes.io/application: wordpress
spec:
  ports:
    - port: 80
  selector:
    app: wordpress
    component: wordpress-svc
  type: LoadBalancer
---
apiVersion: v1
kind: Service
metadata:
  name: wordpress-hsvc
  labels:
    app: wordpress
    component: wordpress-hsvc
  annotations:
      kubernetes.io/application: wordpress
spec:
  ports:
    - port: 3306
  selector:
    app: wordpress
    component: wordpress-hsvc
  clusterIP: None
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: wordpress-webserver
  labels:
    app: wordpress
    component: wordpress-webserver
  annotations:
      kubernetes.io/application: wordpress
spec:
  replicas: 1
  service: wordpress-hsvc
  selector:
    matchLabels:
      app: wordpress
      component: wordpress-webserver
  template:
    metadata:
      labels:
        app: wordpress
        tier: frontend
    spec:
      containers:
      - image: wordpress:4.8-apache
        name: wordpress
        env:
        - name: WORDPRESS_DB_HOST
          value: wordpress-mysql
        - name: WORDPRESS_DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: mysql-pass
              key: password
        ports:
        - containerPort: 80
          name: wordpress
        volumeMounts:
        - name: wordpress-persistent-storage
          mountPath: /var/www/html
  volumeClaimTemplates:
       - metadata:
           name: mysql-persistent-storage
         spec:
           accessModes: [ "ReadWriteOnce" ]
           resources:
             requests:
               storage: 250Gi
```

Above we generate two Applications, one for the MySQL RDBMs and another for the WordPress web server. Each of these 
Applications can be managed independently. For instance, if the WordPress web server is deleted, the components of the 
MySQL RDBMS will not be affected. However, the dependency between the two workloads is made explicit in the 
`Spec.Dependencies` of the `wordpress` Application. Tools and UIs that examine the `wordpress` application are made 
aware of its dependency on the `mysql` Application, and they can surface the dependency to users.

#### Selecting Components
In order to find the `Spec.Componenets` of an Application, a client should perform the following.

1. For each GroupVersionKind in the Applications `Spec.Components`, list all objects of that Kind, in the namespace of 
the Application, that match the Application's `Spec.Selector`.
1. For each selected object, validate that the object has a `kubernetes.io/application` that matches the name of the 
Application. Discard any objects that do not have such an annotation. This ensures that selector overlap does not 
violate the intentions of the application creators and administrators.
1. Unless otherwise indicated, clients should ensure that the `OwnerReferences` of the object contain the `UID` of the 
Application.

#### Retrieving Applications
In order to retrieve the Application corresponding to an API object, a client should perform the following.

1. Examine the object for a `kubernetes.io/application` annotation.
1. If no annotation is found, the object is not associated with an Application.
1. If the annotation is found, retrieve the Application object with the indicated name from the same namespace of the 
object.
1. If no such Application is found, the Application is either deleted or has not yet been created.
1. If the Application is found, validate that the `Spec.Selector` of the Application object matches the labels of the 
object. If the object's labels do not match, the object has been orphaned.
1. Check the `OwnerReferences` of the object for a reference to the Application by `UID` to ensure that ownership has 
been established.

#### Retrieving Dependencies
In order to retrieve the `Spec.Dependencies` of an Application, a client should perform the following.

1. For each name in the `Spec.Dependencies` list, retrieve the Application of that name, in the namespace of the 
Application under consideration. If no such Application is found, the dependency is not installed.

#### Health Checks
The semantics of readiness for an Application vary too greatly for declarative specification. We propose to encode 
the readiness semantics of Applications in the same manner as Pods (i.e. by using programs). An ApplicationHealthCheck 
object contains the template for a Pod that encodes a semantic readiness check for the components of an Application. 
The requirements of implementing a ApplicationHealthCheck Pod are as follows.

1. A health check Pod must implement a readiness check.
1. The readiness check of the Pod may only return successfully when the conditions of readiness are met for the entire 
application. 
1. An ApplicationHealthCheck Pod should not supply other application telemetry. When its readiness check fails that 
telemetry would be unavailable (at precisely the time it is most necessary).

### Application Controller
The Application Controller will be responsible for configuring ownership to enable GC (Garbage Collection) for an 
Application's components and for updating the `Status` of the Application with respect to installation 
status of the components and dependencies of the Application and to the readiness of the Application's health check.

#### Garbage Collection
In order to facilitate GC for Application components, an  `OwnerRefernce` from the API object must be 
added to all of its indicated component objects. As an Application can be composed of arbitrary Kinds, like the 
Garbage Collector, the Application Controller will have to use the discovery API to determine when new Kinds are 
registered with the API server. It will have to watch all Kinds that may be included in an Application's `Components`.

When the Application Controller detects a newly created or updated object it shall do the following.
1. [Retrieve the Application for the object](#retrieving-applications).
1. If the Application's `Selector` matches the object, and if the object does not contain an `OwnerReference` to the 
Application, the Application Controller shall add an  `OwnerReference` from the Application to the object.

As there is no guarantee that an Application's components will be created prior to the Application, the Application 
Controller will do the following upon creation of an Application.

1. [Select the Application's components](#selecting-components).
1. For all selected objects, the Application Controller will add an `OwnerReference` from the Application to the object.

When an Application object is deleted, the Garbage Collector will detect that the owner of the Application's components 
no longer exists at the API Server. This will trigger deletion of the children objects. Note that the Application 
Controller does not establish an ownership relationship between an Application and its dependencies. This is by design.
Also note that, if the Application is composed of Kubernetes primitives that, in turn, aggregate other objects, it is 
sufficient to only add an OwnerReference to the top level object. For instance, deleting a Application containing a 
Deployment will delete the Deployment's ReplicaSets and Pods.

#### Installation Status Reporting
The Application Controller will additionally update the `Status` of Application objects. When the Application Controller 
detects newly created objects, for all objects other than Applications, it shall do the following.

1. [Retrieve the Application for the object](#retrieving-applications).
1. If the Application's `Selector` matches the object, the Application Controller shall update the Applications status 
to indicate that component is installed.

For all Application objects, the Application Controller shall do the following upon creation.

1. [Select the Application's components](#selecting-components).
1. For each selected object, the Application Controller will update the Application object's `Status` to indicate that 
the component is installed. If a component is not found, the Applications status will be updated to indicate that the 
component is not installed.
1. [Retrieve the Application's dependencies](#retrieving-dependencies).
1. For each dependency, the Application controller will update the the Application object's `Status` to indicate the 
installation status of the dependency.
1. The Application Controller will then scan all installed Applications in the same namespace as the Application 
under consideration. If any of those Applications list the Application under consideration in their `Spec.Dependencies`, 
their `Status` will be updated to indicate that the dependency has been installed.

For all objects other than Applications, upon object deletion, the Application Controller shall do the following. 

1. [Retrieve the Application for the object](#retrieving-applications).
1. Update the Application's `Status` to indicate that component is no longer installed.

For all Application objects, upon deletion, the Application Controller will do the following.

1. List all Applications in the same namespace as the application under consideration.
1. For all Applications that contain a reference to the Application under consideration in their `Depepndecies`, the 
Application Controller will update the Application's `Statuts` to indicate that the Application under consideration is 
no longer installed.

For all objects other than Applications, when an object is updated, the Application Controller will do the following.

1. [Retrieve the Application for the object](#retrieving-applications).
1. If the Application's `Selector` does not match the object's labels, the Application Controller will update the 
Applications `Status` to indicate that component is no longer installed and remove any `OwnerReference` from the 
Application to the object.

#### Health Check Pods
If the Application Controller observes the presence of an ApplicationHealthCheck object in the `Spec` of an 
Application, it will do the following.

1. Generate a unique name for a Pod based on the hash of the `Template` field of the HealthCheck. The name generation 
method should be analogous to that of ReplicaSet and Deployment.
1. Create a new Pod based on the `Template`.
1. The Application Controller shall delete and recreate HealthCheck Pods upon mutation of `Spec.HealthCheck` field of 
the application.

#### Health Status Reporting.
The Application Controller will do the following to update the `Stauts.Ready` ConditionStatus of an Application object.

1. If the Application Controller observes an Application object with a nil `Spec.HealthCheck`, it shall set the 
`Status.Ready` field of the Application object to `UNKNOWN`.
1. If the Application Controller observes a Pod created from an ApplicationHealthCheck it shall 
[retrieve the Application for the object](#retrieving-applications).
   1. If the Pod has no Ready Condition, the Application Controller shall set the `Status.Ready` field of the Application 
   object to `UKNOWN`.
   1. If the Pod has a Ready Condition, the Application Controller shall set the `Status.Ready` field of the Application 
   object to the Status of the Pod's Ready Condition.

#### Generation Observation
Each time the Application Controller observes the creation or mutation of an Application object, it will set the 
`Status.ObservedGeneration` to the `Generation` of the Application. As monotonicity of the object's `Generation` is 
ensured by the API Server, the `Satatus.ObservedGeneration` can be used to determine if the Application Controller 
has observed a mutation an Application object.

#### Adopting Existing Components
The design of [Garbage Collection](#garabage-collection) allows existing objects to be aggregated into an Application. 
In order to achieve this users should do the following.

1. Create an Application object that contains the objects that will be adopted in the Application's `Spec.Components`.
1. Annotate the objects with a `kubernetes.io/application` label that references the Application.

The Application Controller will allow the newly created Application to adopt the existing object. By design, this 
workflow should be easily achievable programmatically. Existing tools may wish to implement a command that creates 
an Application in this manner.

### kubectl

Initially, kubectl will support following commands for the Application object.

1. kubectl create
1. kubectl get 
1. kubectl apply
1. kubectl delete
1. kubectl describe

These commands represent basic CRUD operations, and given the state of the discovery API there should be no special 
consideration for their implementation.

### UI Considerations
This section contains information for UI designers on how to use the Application object to surface application 
information to the end user.

#### Discovering Application Components and Dependencies
In order to discover the components and dependencies of an Application, UI designers can do the following.

1. [Select the Application's components](#selecting-components). Note that the Application's components may also 
aggregate sub-components. For instance, a Deployment aggregates ReplicaSets, and those ReplicaSets aggregate Pods. 
UI designers must decide for themselves the best way to view the hierarchical structure of an application. As UI 
designers are already faced with this challenge for objects like Deployment, the addition of the Application object 
should not present any additional, undue burden.
1. [Retrieve the Application's dependencies](#retrieving-dependencies). Again, designers will have to decide 
how they wish to deal with the hierarchical structure of an aggregated application.

#### Discovering an Object's Application 
1. For objects that have a `kubernetes.io/application` annotation UI designers can 
[retrieve the Application for the object](#retrieving-applications) as described above.

#### Application Status

1. For Applications that provide a [health check](#health-checks), UIs may use the `Status.Ready` condition to 
surface the readiness of the application to users. If an Application has no health check, the `Status.Ready` field 
will always be `UNKNOWN`.
1. UIs may surface the installation status of an Application's components and dependencies by examining the set of 
installed components in the Application's `Status.Installed` field.

### Risks and Mitigations

1. In order for the Application API to be useful, we must have broad adoption. In order to mitigate the risk of 
a lack of adoption, we will gather consensus from developers in SIG Apps prior to promoting the API to Beta.
1. As discussed in [Drawbacks](#drawbacks), our approach to the implementation of the Application Controller may 
increase the CPU and memory footprint of the Controller Manager process. We will mitigate this risk by mirroring the 
efficient implementation of the Garbage Collector.

## Graduation Criteria

1. The graduation criteria from Alpha to Beta is simply the consensus of opinion among SIG Apps and SIG CLI that the 
feature is ready to be consumed by tool and UI developers. We expect developers to prototype against the API at alpha 
level stability, but, as it will be disabled by default, we do not expect to be able to gather feedback from end users.
2. The graduation criteria from Beta to GA is adoption and interoperability. Before graduating the feature to GA we 
would like to see broad adoption by tools (e.g. Helm, Bitnami installer) and UIs (Kubernetes OSS UI, GKE, OpenShift, 
Azure).

## Implementation History

### Planned

- Kubernetes 1.10 - Initial implementation of the Application Kind and Application Controller in the apps/v1alpha 
Group Version.
- Kubernetes 1.11 - Promotion of the Application Kind and Application Controller to the apps/v1beta1 Group Version. The 
Application Controller and its associated API are enabled by default. The Application Kind is deprecated in the 
apps/v1alpha Group Version.
- Kubernetes 1.12 - The Application Kind is removed from the apps/v1aplha1 Group Version.
- Kubernetes 1.13 - The Application Kind is promoted to the apps/v1 Group Version and deprecated in the apps/v1beta1 
Group Version.
- Kubernetes 1.14 - Storage migration is complete, on upgrade, to the apps/v1 Group Version of the Application Kind and 
this version is disabled by default.
- Kuberentes 1.15 - The apps/v1beta1 version of the Application Kind is removed from the API.

## Drawbacks

### Increased Memory, CPU, and Network Requirements
The Application Controller, as designed, like the Garbage Collector, must watch a large fraction of the API objects 
stored on the API server. We must be careful to implement the Application Controller efficiently to minimize 
the increased memory and CPU requirements of the core Controller Manager and to minimize pressure on the API server. 

### Increased API Server Storage Requirements
As the Application object will be stored on the API Server, if the users create many Applications, the API Server 
(i.e. etcd) will have an increased storage requirement. Users already create many ConfigMaps and Secrets. It is worth
pointing out that if Application metadata is stored on every object using label and annotation conventions, the 
storage burden on the API Server will likely be greater than if we aggregate the information in a single object and 
store it once. Also, users are already using CRDs for a similar function, and, for those users, the API Server already 
tolerates the additional storage burden.

## Alternatives 
This section contains various design alternatives that were considered during the generation of this KEP.

### Use Labeling Conventions
Rather than introducing an API object, we could require that all tools use a standard labeling scheme to identify 
the components of an Application. This would also provide an opt-in method for tool interoperability. However, this 
method provides no way to garbage collect Application components, express dependencies between applications, or to 
aggregate application meta-data in a single, well-known object. Also, there is no way to ensure that tools implement 
labeling conventions consistently or correctly. Well intentioned implementors may accidentally produce tools that 
incorrectly implement the convention leading to degraded or broken user experiences for other tools or UIs.
This method lacks the capacity for server side validation of the applications definition. There is not way to 
ensure a collection of labels correctly communicates the users intentions. Lastly, this method lacks any 
method to provide health checks.

### Use Custom Resource Definitions
Rather than introducing an API object, each tool could simply use a CRD (Custom Resource Definition) to represent 
Applications. As CRDs are, by definition, a custom API extension mechanism, it is unlikely that this method would 
provide for the desired level of interoperability across tools and UIs. If each tool uses their own CRD, they will a 
priori not be interoperable, and no UI could hope to surface relevant information for all of them in a uniform manner. 
If all tools use a common CRD, then we might as well introduce that resource as a well known Kind.

### Application Type Aggregation
An Application object must be created for each instance of application and not each type of application. It 
might be useful to provide an object that aggregates Applications by type (e.g. all instances of WordPress that are 
installed in the system), but this can also be accomplished by querying for all Applications and filtering or 
aggregating based on the declared `.Spec.Type` field. In the future, if the Application Kind is widely adopted, it would 
be possible to extend the system with another Kind that provides this aggregation via the API object model. The 
challenges here would be multi-tenancy and namespace isolation. If a higher level object aggregates all Applications 
of the same Type, it would have to do so across namespaces. This is not compatible with the existing namespace model 
for Kubernetes unless we assume that the higher level object is only accessible to cluster administrators with 
permissions to access all namespaces in the cluster.

### Use the Garbage Collector Instead of an Application Controller
We could potentially save core hours and memory by incorporating the control logic of the Application Controller 
directly into the Garbage Collector. This design would have a poor separation of concerns and lead to poor code 
modularity.