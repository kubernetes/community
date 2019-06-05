# Mailing list and Github usage best practice                                                       
                                                                                
The Kubernetes Mailing list or Google Groups functions as the primary means of  
asynchronous communication for the project's                                    
[Special Interest Groups (SIG)][sig-list] and [Working Groups (WG)][wg-list].  
That's why you may want to set your filters in your email account to attain a
good signal-to-noise ratio with regards to the mailing list messages and Github
notifications. All the steps below are basically for Gmail users. 

Note: Alternatively, we highly encourage people to use [Gubernator][gubernator] 
and [improve][repo] it instead of setting these filters.

Note: If you are looking to create and manage lists as a chair or other organizer, 
check this mailing list [guideline][mailinglist-guidelines] doc.

- [Example for setting filters to Kubernets Mailing lists](#example-for-setting-filters-to-kubernets-mailing-lists)
- [Examples for setting filters to Kubernetes Github notifications](#examples-for-setting-filters-to-kubernetes-github-notifications)
                                                                                
## Example for setting filters to Kubernetes Mailing lists           

It depends on the SIG or/and WG you are involved in. 
You can setup filters for your Gmail account to be able to categorize emails from 
different mailing lists.
Create a filter following the procedure below:

- In your Gmail account click on **Settings**:
  - **Filters and Blocked Addresses** -> Scroll down and click **create a new filter**
  - In the **to** fields write the email of the SIG's Google Group. 
  - **Create filter** -> Check the box ** Apply the label** and create new
label by choosing **New label...** in the dropdown list.
  - Click on the **Create filter**.
- You can also use filters directly for lists.
  - Matches: list:"kubernetes-dev@googlegroups.com"
  - Do this: Apply label "lists/kubernetes-dev"

## Examples for setting filters to Kubernetes Github notifications

In order to not get spammed by Github notifications and to get all the 
notifications in the right place, here is an example of filters to create/apply
on your Gmail account:

- Stick a blue label on anything kubernetes-related:
  - Matches: (kubernetes OR kubernetes-client OR kubernetes-sigs OR kubernetes-incubator
OR kubernetes-csi)
  - Do this: Apply label "k8s", Mark it as important
- Archive your own actions (sending these is an option in Github's settings).  
You can send them but also archive them, so whenever you need to see the history of 
an issue you can:
  - Matches: to:(your_activity@noreply.github.com) 
    - Do this: Skip Inbox, Mark as read
- Skip bot comments:
  - Matches: (from:(notifications@github.com) (from:(k8s-merge-robot) OR 
from:(Kubernetes Prow Robot) OR from:(k8s-ci-robot)))
    - Do this: Skip Inbox, Mark as read
- Skip push notifications:
  - Matches: to:(push@noreply.github.com) 
    - Do this: Skip Inbox, Mark as read 
- Red label on things assigned to you and/or things request to be reviewed:
  - Matches: to:(assign@noreply.github.com)
    - Do this: Star it, Apply label "gh/assigned", Mark it as important 
  - Matches: to:(review_requested@noreply.github.com)
    - Do this: Star it, Apply label "gh/requested_review", Mark it as important
- Orange label on things you commented on:
  - Matches: to:(comment@noreply.github.com) 
    - Do this: Star it, Apply label "gh/commented" 
- Yellow label on things you have been mentioned on:
  - Matches: to:(mention@noreply.github.com) 
    - Do this: Apply label "gh/mentioned" 
- Grey label:
  - Matches: to:(team_mention@noreply.github.com)
    - Do this: Apply label "gh/team_mention"
  - Matches: to:(author@noreply.github.com) 
    - Do this: Star it, Apply label "gh/authored", Mark it as important 
- Skip messages about issues that you are not participating in, but leave them unread:
  - Matches: from:(notifications@github.com) to:(subscribed@noreply.github.com) 
    - Do this: Skip Inbox 
- Categorize per repository:
  - Matches: list:(community.kubernetes.github.com)
    - Do this: Apply label "k8s/community"   

This [thread](https://groups.google.com/forum/#!topic/kubernetes-dev/5qU8irU7_tE/discussion) in kubernetes-dev google group
is also useful for setting up filters in Gmail.
 
[sig-list]: /sig-list.md#master-sig-list
[wg-list]: /sig-list.md#master-working-group-list
[mailinglist-guidelines]: /communication/mailing-list-guidelines.md
[gubernator]: https://gubernator.k8s.io/pr 
[repo]: https://git.k8s.io/test-infra/gubernator 
