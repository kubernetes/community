# Dynamic Resource/Shortcut Fecthing

Status: Pending

Implementation Owner: @shiywang

## Motivation

Previously, when a new resource added to server, developer should add those resources/shortcuts name into two places of kubectl: [here](https://github.com/kubernetes/kubernetes/blob/357db0c39c7203809b369516c0ec93831ead8649/pkg/kubectl/cmd/cmd.go#L201) and [there](https://github.com/kubernetes/kubernetes/blob/f9913c4948038b9b51f2342134d546c6bb74e7a3/pkg/kubectl/kubectl.go#L46). then `kubectl get/describe/explain... -h` will show the new added hard code resources/shortcuts, But some times people always forget to add those things, and we have to submit new pr to fix that everytime.

## Proposal

We should eliminate those every time add new hardcoded resources in kubectl, improve development experiences, but for how to prompt those information properly, there's two solutions here I've come up with: 

### solution 1

Keep current help messages, and add a ring1Factory function use clientDiscovery fetch resources/shortcuts from server then concatenate those valid resources string and print every time when user execute command like this:

```
➜  ~ kubectl get -h
Display one or many resources. 

Valid resource types include: 

  * all  
  * certificatesigningrequests (aka 'csr')  
  * clusterrolebindings  
  * clusterroles  
  * clusters (valid only for federation apiservers)  
  * componentstatuses (aka 'cs')  
  * configmaps (aka 'cm')  
  ...
```

### solution 2

Remove those valid resources list in help messages, and add a dedicated command to fetch resources/shortcuts, like `kubectl api-resources` similiar to `kubectl api-versions` or we could extend `api-versions` to support that, then tell user use `kubectl api-resources` to check what kind of valid resouces they can use.

```
➜  ~ kubectl get -h
Display one or many resources. 

You can use `kubectl api-resources` to check valid resource types server provided. 

...
```


## User Experience

### Use Cases

`kubectl get/describe/explain..   -h`


### Shortage

solution 1 may have some shortage in first time prompt help messages.

    ResourceAliasesResourceShortFormForkubectl get -h  0.13s user 0.02s system 128% cpu 0.117 total
    ./kubectl get -h  0.13s user 0.02s system 5% cpu 2.600 total



## Implementation

### solution 1

Add function/interface in ring1Factory  `func ValidResourcesFromDiscoveryClient []ResourceShortcut`

change for resource in kubectl get/describe/explain... -h

    pkg/kubectl/cmd/get.go
    cmd := &cobra.Command{
        ...
        Short:   i18n.T("Display one or many resources"),
        Long:    templates.LongDynamicDesc(getLong, f.ValidResourcesFromDiscoveryClient())  
        Example: getExample,

    pkg/kubectl/cmd/templates/normalizers.go    
    func LongDynamicDesc(s string, r []kubectl.ResourceShortcuts) {
    	if len(r) != 0 {
    		//foreach ResourcesShortcuts sorted, and concatenate to string  
    		return LongDesc(fmt.Sprintf(getLong, concatenateString))
    	}
    	return LongDesc(fmt.Sprintf(getLong, validResours))
    }


### solution 2 

This is a simple one, only change is remove hardcoded list in help messages of `kubectl get/explain/describe ...` and implement a new command `kubectl api-resources`.


## Client/Server Backwards/Forwards compatibility

Depend on which solution we choose.

## Other things need to discuss

1. Use binary cache from OpenAPI instead of cache from `.kube/cache/discovery` ?

2. Also there's another two functions in `pkg/kubectl/kubectl.go`: `ResourceShortFormFor` and `ResourceAliases`, we may also want to replace it with dynamic fetching, change signature of those two functions and also fetch resources from discoveryClient, this work is a follow up of PRs: https://github.com/kubernetes/kubernetes/pull/40312 and https://github.com/kubernetes/kubernetes/pull/38835 I think we have no reason to not eliminate those hardcoded resources/shortcuts.




