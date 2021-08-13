# Meeting notes

- CCM per cloud provider
- same host as kube-apiserver
- caches live in memory
- refresh cache, but can be forced to by request
- Controller manager attempts to use PoLA, but the service account controller has permission to write to it's own policies
- Cloud controller (routes, IPAM, &c.) can talk to external resources
- CCM/KCM have no notion of multi-tenant, and there are implications going forward
- Deployments across namespace
- cloud controller has access to cloud credentials (passed in by various means, as we saw in the code)
- CCM is a reference implementation, meant to separate out other company's code
  - So Amazon doesn't need to have Red Hat's code running, &c.
- shared acache across all controllers
- [FINDING] separate out high privileged controllers from lower privileged ones, so there's no confused deputy
  - single binary for controller
  - if you can trick the service account controller into granting access to things you shouldn't (for example) that would be problematic
  - make a "privileged controller manager" which bundles high and low-privileged controllers, and adds another trust boundary
