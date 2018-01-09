Kubernetes Dashboard UX breakout session 12.5.17, led by Rahul Dhide ([rahuldhide](https://github.com/rahuldhide)) and Dan Romlein

* **Resources:**

    * Dashboard [User Types #975](https://github.com/kubernetes/dashboard/issues/975) 

    * Dashboard [User Types and Use Cases](https://docs.google.com/document/d/1urAlgRP7AbcdsOMQ_piQQ6O1XTDIum_LmOUe8xsC4pE/edit)

    * [SIG UI weekly](/sig-ui)

* **Notes**

    * 2018 Dashboard strategy.

        * [Deck](https://docs.google.com/presentation/d/1q1G1vWCrenI3GVsyF4d-2rpyVrdJgFW7fIbvcAAd9Lg/edit?ts=5a26f250)

        * [Github issue](https://github.com/kubernetes/dashboard/issues/2556)

    * **Kubectl access via Dashboard**.

        * Dhaval: Provide context to issues.

    * **Third-party Widgets**.

    * **Custom Views**.

        * Dhaval: Custom views will be very useful to share the event details, contextual information, and logs with specific time ranges. It will help our development teams to quickly analyze the issues. 

        * Henning: UI should allow users to pick the important properties in different views/widgets.

    * **Integrations and CRDs**.

    * **Onboarding experience**.

        * Jared: Currently there is no crosslinking between docs and the UI. We can explore the concept and impact further.

    * **Design system.** 

        * e.g. Standard for how a pod’s status is displayed.

    * Different use cases/personas

        * application developer

        * application operator

        * (multi-)cluster operator

    * **Feedback from Dhaval and Hennings:**

        * "We don’t use Dashboard, because of lack of authorization control."

        * "Dashboard today caters to the dev perspective, which is OK."

            * But for ops, currently missing dependency stacks (e.g. OS version) "I want to know different node versions."

        * "Our developers use Dashboard for the logs"

        * "When we perform troubleshooting & debugging we expect 30 min before and after incident. 

        * Want to be able to link users to docs for more info; "What’s a pod? What’s a deployment?"

            * *Contextual docs* displayed in UI. 

            * Idea: Dashboard could scrape docs. 

        * Kubernetes docs working on expanding [glossary](https://kubernetes.io/docs/reference/glossary/?fundamental=true)

        * Dashboard is backwards compatible. 

        * Demo of [Kubernetes Operational View](https://github.com/hjacobs/kube-ops-view)

        * Use Case:

            * Onboarding: explain resource limits vs. resource requests. 

            * Quickly looking at a cluster and knowing what’s going on.

        * "I really like this UI"

        * Wanted to gamify K8s with this UI. 

        * "My cluster has 118 nodes, so the ability to filter would be important." (general scale issues)

        * "Problem we’re running into is that the number of nodes will crash the browser, so we need some way to select cluster first"

        * Defining a view is hard, Dashboard shouldn’t attempt to do that.

        * Use case for custom Dashboard re-skinning.

            * "My exec looks at the look and feel."

* Kubernetes Dashboard 2017 User Survey: [https://docs.google.com/forms/d/e/1FAIpQLScnxeub_xh7Lp4iZO1RKdTgIYK_cTwqFKv1WD-Cue4tZHcbhw/viewform?usp=sf_link](https://docs.google.com/forms/d/e/1FAIpQLScnxeub_xh7Lp4iZO1RKdTgIYK_cTwqFKv1WD-Cue4tZHcbhw/viewform?usp=sf_link)

