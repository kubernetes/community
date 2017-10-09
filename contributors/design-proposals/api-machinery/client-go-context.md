# Using context.Context with client-go

* Authors: @maleck13

## Abstract

This proposal aims to outline how to allow using the [context package](https://golang.org/pkg/context/) with [client-go](https://github.com/kubernetes/client-go) to allow propagation of cancellation and timeouts 
and also to allow exposing context based http tracing [go blog on context based httptrace](https://blog.golang.org/http-tracing) to consumers of the client.

## Motivation and Goals

When using client-go, external calls to the Kubernetes API are made in order to manage resources. The initiator of these calls may find due to an external event, 
that they no longer need the resource they had requested, but they have no means by which to inform the client of this change. 

When using client-go, having the ability to add context based http tracing would be valuable for issues around performance, observability and debugging. 

- Allow consumers of the client to indicate, in an idiomatic way, that the action that caused them to invoke the client has been cancelled.
- Allow consumers of the client to instrument and gain insight via context based http tracing.

## Non Goals

- Introduce a transactional system for write operations whereby upon cancellation the client would somehow rollback any current write operations.

## API

There are two approaches outlined. The first is a non breaking change that appends to the existing API for resources but perhaps puts more burden on the client consumer. 
The second is a breaking change to the client-go api, that is potentially more idiomatic and clear.
 
Common to both is the fact that the request type in client-go  already exposes a ```Context```  function: [request context set on client-go rest.Request](https://github.com/kubernetes/client-go/blob/master/rest/request.go#L393). 
If set, this context is passed all the way through to the underlying ```http.Request```: [request context passed to http request code](https://github.com/kubernetes/client-go/blob/master/rest/request.go#L484). 
So to achieve the goals, we need only modify the client API to allow passing in an external context that is then used to set the existing context on the request type.

### Non breaking change

The client resources (Pods, Secrets etc...) each expose a resource specific interface ```PodInterface``` for example. This is returned from a ```Getter``` interface:

```go
type PodsGetter interface {
	Pods(namespace string) PodInterface
}
``` 

If we were to modify these resource interfaces and their underlying concrete types ```pod``` in this case, to add a ```WithContext``` method, it would allow consumers of the client to set context while maintaining backwards compatibility for all other consumers of the API.

```go
// PodInterface has methods to work with Pod resources.
type PodInterface interface {
  	WithContext(ctx context.Context) PodInterface
    ...
}

type pods struct {
	client rest.Interface
	ns     string
	ctx    context.Context
}

// WithContext allows you to set a context that will be used by the underlying http request
func (c *pods) WithContext(ctx context.Context) PodInterface {
  	c.ctx = ctx
  	return c
}

```

To pass through this context, it would be necessary to change the underlying client calls to accept the context. Example:

```go 
func (c *pods) Get(name string, options meta_v1.GetOptions) (result *v1.Pod, err error) {
	result = &v1.Pod{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pods").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}
```

This would result in an API that would be interacted with like so if the consumer wished to pass a context:
```go  
   ctx := req.Context()
   pod, err := k8client.CoreV1().Pods(namespace).WithContext(ctx).Get(podName, ...)
```


HTTP Tracing and cancellation are also likely useful when using ```informers``` and ```listers```, as these also use the client interfaces backwards compatibility can be maintained by adding a new set of constructors for adding a
specific context object:

```go
func NewConfigMapInformerWithContext(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, ctx context.Context) cache.SharedIndexInformer {}

func (f *configMapInformer) InformerWithContext(ctx context.Context) cache.SharedIndexInformer {}

func defaultConfigMapInformerWithContext(client kubernetes.Interface, resyncPeriod time.Duration, ctx context.Context) cache.SharedIndexInformer {
	return NewConfigMapInformerWithContext(client, meta_v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},ctx)
}

func (f *configMapInformer) ListerWithContext(ctx context.Context) v1.ConfigMapLister {
	return v1.NewConfigMapLister(f.InformerWithContext(ctx).GetIndexer())
}
```

**Concerns / Notes**
- The idea behind using a context with a request is it is only meant to live as long as the initial request that created it. If we expose an API like the one outlined, the consumer would need to know that they needed to call
```WithContext(ctx context.Context)``` with each request context.

- The fakes would need to be updated and new tests added to cover the new APIs.

- The number of entry points and maintenance points is increased. 

### Breaking change

The second option is to pass the context with each action on a resource. This may be more idiomatic and explicit: 

```go 
func (c *pods) Get(ctx context.Context, name string, options meta_v1.GetOptions) (result *v1.Pod, err error) {
	result = &v1.Pod{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pods").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Context(c.ctx).
		Do().
		Into(result)
	return
}
```

These change would impact on every API. Although the changes required would be many, the simplest approach to maintaining current behaviour would be to pass in nil or ```context.ToDo()```.

For the ```listers``` and ```informers``` the change outlined above (adding new constructors) would still allow a context to be passed through to client functions doing the work.

**Concerns**

This is very clearly a breaking change that would likely require extensive changes to both tests and implementation code throughout the impacted code bases.

