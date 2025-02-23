# SIG Node Meeting Notes

## Jan 14, 2025

* \[Tim Allclair\] Changing the CRI contract for `UpdateContainerResources`, to require it to not intentionally restart a container. Redefine (and probably rename) the "NotRequired" resize restart policy to match, meaning containers won't be deliberately restarted (could still trigger an OOM). Runtimes that can't support in-place resize should return an error (e.g. VM runtimes).

## Jan 7, 2025

* \[Tim Allclair\] Shrinking memory limits bellow memory usage (in-place resize)  
  * See [In Place Pod Resize - Handling Memory Limit Decreases](https://docs.google.com/document/d/1cEFLXKwNOSNLAkzyhoJUgkBW0OiX-9bXB_aJV7OAypw/edit?tab=t.0)  
    * \[sergey\] to not interrupt the flow posting some questions:  
      * should we disable resize on cgroupv1 to avoid diff behavior?  
        * Mrunal: true, but we like the v1 behavior better  
      * How VPA can use the feature safely? This is another example of VPA cannot guarantee non-disruptive pod resize (best effort)  
        * Tim to follow up with VPA folks  
      * All checks suggested still have possibility for the race condition, the user experience will be super annoying. Should we reconsider the application of memory.high?  
  * Mrunal: Ask Waiman (Red Hat memory subsystem maintainer) if switching to v1 behavior is possible  
    * Peter: we probably need to do this to fully avoid TOCTOU / OOM kills  
  * Tim: Pause container to avoid TOCTOU?  
* \[Tim Allclair\] Pod generation, status generation (design proposal)  
  * See proposal section of[\[PUBLIC\] Design proposal: Pod ResizeStatus Improvements](https://docs.google.com/document/d/10m0vdWbjqF_q_f1_N3gVOD5YeNn09BwFcVVycG_MoEY/edit?tab=t.0#heading=h.3rtsrq1ip5xv)  
* \[Zeel Patel\] [https://github.com/kubernetes/kubernetes/issues/129447](https://github.com/kubernetes/kubernetes/issues/129447)   
  * Create a KEP PR: [https://github.com/kubernetes/enhancements/pull/5022](https://github.com/kubernetes/enhancements/pull/5022)   
  * \[fromani\] I’d be happy to review  
* \[HirazawaUi\]: Addressing the Issue of Slow Pod State Transitions When Evented PLEG is Enabled  
  * Due to the time zone difference, I regret that I am unable to attend the meeting. I have compiled the entire process of how @hshiina and I attempted to address this issue, as well as my proposal, into a detailed document: [https://docs.google.com/document/d/1TPrY56q9MNW8r1FuzKDFkBBhOjQ0hqi7wJAbIP1O-4g](https://docs.google.com/document/d/1TPrY56q9MNW8r1FuzKDFkBBhOjQ0hqi7wJAbIP1O-4g). I sincerely hope this can be discussed during the meeting.  
* \[SergeyKanzhelev\] [Kubernetes feature development and Container runtimes](https://docs.google.com/document/d/1y42XrUPrm-6DZby1RQjexYYoNn822IRR6igsOiy_62c/edit?tab=t.0#heading=h.2bbangbf1ha4)  
  * State containderd and cri-o explicitly  
* \[kannon92\] [Deprecation of NodeFeature in sig-node testing](https://github.com/kubernetes/kubernetes/pull/129166)  
* \[marquiz\] KEP-4112 Pass down resources to CRI, request for review:  
  [https://github.com/kubernetes/enhancements/pull/4113](https://github.com/kubernetes/enhancements/pull/4113)  
* 1.33 time is coming\! [https://github.com/kubernetes/sig-release/pull/2706/](https://github.com/kubernetes/sig-release/pull/2706/files)
