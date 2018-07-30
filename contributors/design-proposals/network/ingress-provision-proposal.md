# IngressClass and LoadBalancer provision proposal

***Authors*** : @mqliang @m1093782566 @haibinxie

## Motivation 

Many Kubernetes clusters have a variety of LoadBalancers which differ widely in 
performance, and other characteristics, especially in an environment such as GCE and AWS,
which has lots of degrees of freedom in API(a single instance of a load balancer, or a 
more complicated setup of frontends that provide GSLB, DDoS protection, etc). 

Administrators need a way to dynamically provision LB of these different classes to 
automatically meet user demand.

A mechanism called “ingress class” is proposed to provide this capability, which allows 
configuration of different LB and supports provisioning multiple types of LB within a 
single cloud. The mechanism already exists in a half-baked form via the `ingress.class` 
annotation.

***NOTE***: in this proposal, “LB” means LoadBalancer that handles Ingress
resource.

In conclusion, the ideal relationship is:

* 1 ingress controller = 1 provisioner
* N classes can reference 1 provisioner, and at least 1 LB for each class
* M Ingresses can reference 1 class

### Out-of-Tree Provisioners

One of our goals is to enable administrators to configure out-of-tree ingress 
provisioners, whose code does not live in the Kubernetes project, but in a 
variety of ingress-controllers. That’s because different platforms have
different attributes on LoadBalancer, for example:

* On AWS/GCP/GCE, LB has a limit on the number of URL map per LB, e.g 50 on GCE.
* On Aliyun, a single LB also have a upper limit of bandwidth (0.5G) for
incoming request.

So, it is impossible to anticipate every aspect and manner of provisioning 
that administrators will want to perform, It’s more reasonable to let ingress
controllers to implement provision logic by themself.

## API Definition

```
// IngressClass describes a named "class" of load balancer offered in a cluster.
// Different classes might map to quality-of-service levels, or to backend implementation,
// or to arbitrary policies determined by the cluster administrators.  Kubernetes
// itself is unopinionated about what classes represent.  This concept is sometimes
// called "profiles" in other systems.
// The name of a IngressClass object is significant, and is how users can request a particular 
//class.
type IngressClass struct {
    metav1.TypeMeta
	// +optional
	metav1.ObjectMeta

	// provisioner is the driver expected to handle this IngressClass.
	// This is an optionally-prefixed name, like a label key. 
    // The prefix name of provisioner can be “kubernetes.io” or “cloud.google.com”,
    // or maybe something else.
	// For example, if the prefix name is “kubernetes.io”, then "kubernetes.io/gce" 
    // will target the GCE controller, forcing the nginx controller to ignore it.
    // "kubernetes.io/nginx" will target the NGINX controller, forcing the GCE
    // controller to ignore it.
	// This value may not be empty.
	Provisioner string

	// parameters holds parameters for the provisioner.
	// These values are opaque to the  system and are passed directly
	// to the provisioner.  The only validation done on keys is that they are
	// not empty.  The maximum number of parameters is
	// 512, with a cumulative max size of 256K
	// +optional
	Parameters map[string]string
}

// IngressClassList is a collection of ingress classes.
type IngressClassList struct {
	metav1.TypeMeta
	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ListMeta

	// Items is the list of IngressClasses
	Items []IngressClass
}
```

#### Open Questions:
* What’s the best prefix name of IngressClass.Provisioner? “kubernetes.io” or
“cloud.google.com”, or maybe something else?

We will  add a new field “IngressClassName” in IngressSpec, which references to the 
ingress class name.

```
Type IngressSpec struct {
            // IngressClassName references to the ingress class name.
            // +optional
            IngressClassName *string
}
```

***NOTE:***

* IngressClass is a cluster-scoped resource.

* The Provisioner field of IngressClass behaves very similar to the current 
“kubernetes.io/ingress.class” annotation. Admission controller will set a default class
if users create an Ingress without specifying the IngressClassName.  If failed to choose
a default class, deploying multiple ingress controllers and not specifying the 
IngressClassName will result in NO controllers touching the Ingress. By the way, 
it’s different from what the current “kubernetes.io/ingress.class” annotation 
behaves - if  “kubernetes.io/ingress.class” annotation set to empty, multiple ingress
controllers will fight satisfying the ingress. However, the existing 
“kubernetes.io/ingress.class” annotation is ad hoc and never really designed. 
“kubernetes.io/ingress.class” annotation will be deprecated in the future.

* There is no RecyclePolicy field in the API definition, because if one LB has no Ingresses
referencing it, it can be removed. We don't need a policy for that - it already works in 
every Ingress controller.

* Currently the only validation for IngressClass.Parameters is checking if string is 
empty. Should we introduce the OpenAPI validation? Probably, but we should do it as a
2nd step if really needed.

## Default Ingress Class

If IngressClassName is not specified in the Ingress, the default ingress class will be 
used for provisioning.

### Install Default Ingress Classes

To reduce the burden of setting up default IngressClasses in a cluster, we can install 
default ingress classes via the add-on manager or by hand. To use these default 
IngressClasses, users do not need to specify IngressClassName in the Ingress.

For GCE, a default ingress class might look like:

```
apiVersion: network.k8s.io/v1
kind: IngressClass
metadata:
  name: standard
  annotations:
    ingressclass.kubernetes.io/is-default-class: "true"
provisioner: kubernetes.io/gce
parameters:
  type: standard-https
```

For bare-metal, a default ingress class might look like:

```
apiVersion: network.k8s.io/v1
kind: IngressClass
metadata:
 namespace: default
 name: standard
 annotations:
   ingressclass.kubernetes.io/is-default-class: "true"
provisioner: kubernetes.io/nginx
```

From the view of implementation, the default IngressClass has an annotation 
“ingressclass.kubernetes.io/is-default-class” set to true. Any other value or absence of 
the annotation is interpreted as false. To mark a IngressClass as non-default, users
should change its value to false.

### DefaultIngressClass Admission Controller

Introducing a new plugin in admission controller - the DefaultIngressClass admission 
controller, which will automatically adds the IngressClassName pointing to the default 
Ingress Class, in case the user did not provide a class name value when creating Ingress.
The new DefaultStorageClass admission controller works only if a default IngressClass
exists. The controller logic is:

* Find available IngressClasses,
* Figure which is the default,
* Write the default class to incoming Ingress objects.
* If No default class selected, admission controller will do nothing about the Ingress.

## Kubectl Output Format

### List the IngressClasses in the cluster:

```
$ kubectl get ingressclass
```

The output format is [Ingress class name], [Provisioner type], which is similar to this:

```
NAME                    TYPE
silver(default)         kubernetes.io/gce
gold                        kubernetes.io/gce
bronze                   kubernetes.io/nginx
```

The default IngressClass is marked by (default). Both the silver and gold ingress classes 
are back-ended by gce, while the bronze ingress class is back-ended by nginx.

### Show the details of given IngressClass:

```
$ kubectl describe ingressclass standard
```

The output is similar to this:
```
Name:     	    standard
IsDefaultClass:  Yes
Annotations:      storageclass.beta.kubernetes.io/is-default-class=true
Provisioner:       kubernetes.io/gce
Parameters:      type=standard-https
Events:             <none>
```

## Provision implementation examples

Although the provision logic is left to ingress-controller and the implementation is 
totally opaque to user, we still should provide concrete implementation of the provision 
logic for nginx-ingress-controller, envoy(Istio or Contour) and gce-ingress-controller. 
We will describe how the provision logic can be implemented for gce, aws and bare-metal.

Users can manually deploy ingress controllers via kubectl, like

* kubectl create -f gce-ingress-controller.yaml
* kubectl create -f aws-ingress-controller.yaml
* kubectl create -f nginx-ingress-controller.yaml

### GCE implementation:

* user defines 3 classes: Gold, Silver, Bronze, they all use “gce" as the provisioner.

* user creates an Ingress specifying class=Gold

* gce-ingress-controller watch this event, and find “Oh, this Ingress specifying 
class=Gold, and class Gold uses “gce" as the provisioner, I am responsible for it”, 
and begins its work.

* gce-ingress-controller will dynamically provision a new Gold class LB if not exists or 
pick the existing one to serve the ingress rules. Since there is no limitation about rule 
number or incoming request bandwidth for gce LB, so we can use only one Gold-LB to handle
ALL Gold class Ingresses.

### AWS implementation:

* user defines 3 classes: Gold, Silver, Bronze, they all use “aws" as the provisioner.

* user creates an Ingress specifying class=Gold.

* aws-ingress-controller watch this event, and find “Oh, this Ingress specifying 
class=Gold, and class Gold uses “aws" as the provisioner, I am responsible for it”, 
and begins its work. aws-ingress-controller will dynamically provision a new Gold class
LB if not exists or pick the existing one to serve the ingress rules. At some times, 
aws-ingress-controller aware that the rules bound to Gold-LB reaches the upper bound
(remember that on AWS, LB has a limit on the number of rules per LB), it will auto 
provision a new Gold class LB, say Gold-LB-2, to handle other Gold class Ingresses. 

### Nginx Implementation:

* user defines 3 classes: Gold, Silver, Bronze, they all use “nginx" as the provisioner.

* The nginx-ingress-controller will provision only one nginx pod set, and configure nginx 
instance with different config sections for different classes. 

* If you want all the Ingresses share a single IP, you should put only one LVS cluster 
in front of nginx pods. However, if you want different class use different IP, you should 
set up multiple LVS clusters, and put them in front of the corresponding nginx pods. 

And the provision logic is totally opaque to a kubernetes user, all they saw was 
"create an Ingress, get an IP".
