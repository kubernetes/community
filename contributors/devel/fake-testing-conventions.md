Fake Testing Conventions
===============

Updated: 2/25/2017

## Overview

To test non-trivial components of the Kubernetes one needs fake components that simulate behaviour of the corresponding real counterparts. From creating interfaces, choosing the correct place for new data types to creating testing packages. This document attempts to capture good practices of the process.

## Good practices

* If a fake component is used for testing, it should by defined in testing package

```go
# plugin/pkg/scheduler/testing/fake_lister.go
package testing

...
type NodeLister []*v1.Node

// List returns nodes as a []string.
func (f NodeLister) List() ([]*v1.Node, error) {
 return f, nil
}
```

* Choose proper names, e.g. instead of ``FakeNodeLister`` use ``NodeLister``.

```go
import (
   fakescheduler "plugin/pkg/scheduler/testing"
)

...
nodeLister := fakescheduler.NodeLister(nodes)
...
```


* Fake component should implemented the same interface as its real counterpart

```go
# plugin/pkg/scheduler/algorithm/types.go
package algorithm

// NodeLister interface represents anything that can list nodes for a scheduler.
type NodeLister interface {
	// We explicitly return []*v1.Node, instead of v1.NodeList, to avoid
	// performing expensive copies that are unneeded.
	List() ([]*v1.Node, error)
}
```

```go
# plugin/pkg/scheduler/testing/fake_lister.go
package testing

import (
  types "plugin/pkg/scheduler/algorithm"
)

type NodeLister []*v1.Node

// List returns nodes as a []string.
func (f NodeLister) List() ([]*v1.Node, error) {
 return f, nil
}

// Let the compiler check the fake type implements the interface
var _ types.NodeLister = &NodeLister{}
```
