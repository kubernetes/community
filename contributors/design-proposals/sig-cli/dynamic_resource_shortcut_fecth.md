# Dynamic Resource/Shortcut Fecthing

Status: Pending

Implementation Owner: @shiywang

## Motivation

Previously, when a new resources added to server, developer should add those resources/shortcuts name into two places of kubectl: here and there. then kubectl get/describe/explain... -h will show the new added hard code resources/shortcuts, But some times people always forget to add those things, and we have to submit new pr to fix that everytime.

## Proposal

we should fetch those information from server, but for how to prompt those information proply, there's two solutions here: 

### solution 1
keep current help messages, and add an ring1Factory function use clientDiscovery fetch resources/shortcuts from server then concatenet those valid resources string and print every time when user execute command like this:
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

remove those valid resources list in help messages, and add an dedicated command from fetch resources/shortcuts, and tell user use `kubectl new-command` to check what kind of valid resouces they can use.
 ```
➜  ~ kubectl get -h
Display one or many resources. 

You can use `kubectl new-command` to check valid resource types server provided. 

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
ring1Factory change

add function/interface in ring1Factory  func ValidResourcesFromDiscoveryClient []ResourceShortcut

change functions ResourceShortFormFor and ResourceAliases use ValidResourcesFromDiscoveryClient returned values to lookup shortcuts, if len(ResourceShortcut)== 0means fetch from server faild, use  ResourcesShortcutStatic



change for resource in kubectl get/describe/explain... -h

    	cmd := &cobra.Command{
    		...
    		Short:   i18n.T("Display one or many resources"),
    		Long:    templates.LongDynamicDesc(getLong, f.ValidResourcesFromDiscoveryClient()) 
    		Example: getExample,
    		...
    
    func LongDynamicDesc(s string, r []kubectl.ResourceShortcuts) {
    	if len(r) != 0 {
    		//foreach ResourcesShortcuts sorted, and concatenate to string  
    		return LongDesc(fmt.Sprintf(getLong, concatenateString))
    	}
    	return LongDesc(fmt.Sprintf(getLong, validResours))
    }



change for shortcuts in pkg/kubectl/kubectl.go

change signature of those two functions ResourceShortFormFor and ResourceAliases

accept  []kubectl.ResourceShortcuts if exist

### solution 2 
add an new command, and change help messages in `kubectl get/explain/describe ...`


## Client/Server Backwards/Forwards compatibility

### solution 1
Shortname of meta.v1.APIResource was an new added filed in commit https://github.com/kubernetes/kubernetes/pull/40312/commits/d100d5644661c288d52dbe0456a9c7d184f61d31, but theoretically, I think we can do fully backwards/forwards comptibility, need to check with @deads2k.

### solution 2
new kubectl will prompt differently with old server, but I think it won't be a big problem.

## Other things

about fetched resources use binary cache from OpenAPI ?



## Other things you need to know

This work is a follow up of PRs: https://github.com/kubernetes/kubernetes/pull/40312 and https://github.com/kubernetes/kubernetes/pull/38835 I think we have no reason to not eliminate those hardcoded resources/shortcuts.


