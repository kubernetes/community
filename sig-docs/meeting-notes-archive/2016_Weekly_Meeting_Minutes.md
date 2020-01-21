**12/13/2016**

*   1.5 is out the door.  Wooo.
    *   New site org in 1.5
    *   Added to tasks, tutorials, and concepts
    *   Statefulset docs fell under concepts, tutorials, and tasks.
    *   We want to do this the same way going forward.
*   Template issues :(
    *   Github issues with templates :(
    *   They update, we break.  Markdown processor breaks the site.
        *   First update broke the site chrome.
        *   Current change breaks how we automate page titling.
    *   AI: 
        *   Explore Moving to using netlify to stage the site? (Jessfraz, Ddonnelly) 
        *   Fix Page title needs to be fixed. 
        *   Explore fixing templating in github pages.
        *   Propose hiring a web-dev for the team. 
*   TOCs
    *   Some pages that have TOCs in flux (tools page, for example)
    *   Should figure out a standard for organizing these pages. 
        *   Concepts (Ddonnelly)
        *   Tasks, Tutorials (Stevepe)
*   Issues Triaging, Templates, Standards
    *   [https://github.com/kubernetes/kubernetes.github.io/pull/1860](https://github.com/kubernetes/kubernetes.github.io/pull/1860) (AI: jaredb, review) 
    *   Jan 10th is issue triage day :)
*   1.5 RefDocs Update
    *   [https://github.com/kubernetes/kubernetes.github.io/pull/1818](https://github.com/kubernetes/kubernetes.github.io/pull/1818)
    *   Manually reviewing some content, verifying HTML content.
    *   Using Deploy preview to review. 
    *   Mehdy wants to release a blog post around the work that he’s done.
*   Versioning
    *   Some thoughts:
        *   Maintain a release branch for each release
        *   Stage using netlify
        *   Clarify the level of support for release branches
        *   Only 1.5 going forward
*   Deprecation notices
    *   Going to land on the site.  Devin will decide. 

**11/28/2016**



*   Meeting using Zoom
*   1.5 Launches (Ihor, Jared)
    *   Launch Process Updates
    *   Link to spreadsheet https://docs.google.com/spreadsheets/d/1g9JU-67ncE4MHMeKnmslm-JO_aKeltv2kg_Dd6VFmKs/edit#gid=0
*   License Updates (Jared) 
    *   cla/linuxfoundation is required, google license is not.
    *   Associated PRs:
        *   [update LICENSE](https://github.com/kubernetes/kubernetes.github.io/pull/1763)
        *   [update footer](https://github.com/kubernetes/kubernetes.github.io/pull/1764)
        *   [update "how to contribute"](https://github.com/kubernetes/kubernetes.github.io/pull/1765)
*   Style Guides (Steve) 
    *   PRs checked in
    *   Old Templates deleted.  Long live the new templates.
    *   In progress: Remove some content from README and editdocs, and blend it into the new contribution pages.
*   Issues Triaging (Abby)
    *   PR coming soon
    *   1/10 triage day!
    *   Kubernetes swag will be given
*   Quickstart Discussion: Bill Prin
    *   Move getting started and creating a cluster content to “tasks”
    *   Revamp content, remove GKE “Hello Node”
    *   Jeff and Bill will come up with a more in depth proposal for next week. 
*   API Reference docs (Phillip)
    *   [docs.k8s.replicatingperfection.net](http://docs.k8s.replicatingperfection.net)
    *   Questions?
        *   Need folks to review API Docs
        *   Sync with Devin on bandwidth
        *   Soft launch for 1.5? 

**11/15/2016**

Agenda:



*   Site Updates
    *   Issue Triaging (Abby)
        *   [http://kubernetes.io/docs/contribute/](http://kubernetes.io/docs/contribute/)
        *   Should schedule issue-triaging day
        *   Steve suggests review-issue.md
    *   Reference Docs Update? (Phillip)
        *   Should have something in two weeks
        *   Will use the updated refdocs for 1.5
    *   Style and Contributor Guides
        *   PRs ready
    *   Preparing for K8s 1.5 Release: [https://docs.google.com/spreadsheets/d/1PcAW-ZQnzlwEs59ygHdFA-72n6FUcGIc2lx4IeoXvco/edit?ts=57ced759#gid=110968752&vpid=A1](https://docs.google.com/spreadsheets/d/1PcAW-ZQnzlwEs59ygHdFA-72n6FUcGIc2lx4IeoXvco/edit?ts=57ced759#gid=110968752&vpid=A1)
*   Dev Summit Feedback
    *   [Brainstormed Issues with Site](https://github.com/kubernetes/kubernetes.github.io/issues/1686)
    *   [UX issues ](https://docs.google.com/document/d/1BoF_wA7J6N-zuCUgwMMj27ngaj3e4T0YJhvTNwWiWj8/edit#heading=h.y2ttshn5nulr)
    *   [Brainstorm on Developer Onboarding](https://docs.google.com/document/d/19B2vcK6Y3xE3JO7sd4n6lPF0m9bBVZpNucvcR3MwEXA/edit)
*   PR backlog

**10/25/2016**

Agenda



*   Site Updates
*   Issue Backlog
    *   Doc shared with the SIG docs group
    *   [https://docs.google.com/document/d/1j1A6sVW_t5RF_dlxIxpWVbAQeLUlzR1Z26zENBltRms/edit](https://docs.google.com/document/d/1j1A6sVW_t5RF_dlxIxpWVbAQeLUlzR1Z26zENBltRms/edit) 
*   Reference Docs Update
*   Redirects on Github?
*   API and CLI Docs updates
    *   [http://docs.k8s.replicatingperfection.net/](http://docs.k8s.replicatingperfection.net/)
    *   [http://kubectl.k8s.replicatingperfection.net/](http://kubectl.k8s.replicatingperfection.net/)
*   Suggested task docs
    *   

**10/11/2016**

No meeting due to technical issues. 

**9/27/2016**

Announcements



*   1.4 went live!  Hoorays!
*   Templates and new site layout went live! ([http://kubernetes.io/docs/](http://kubernetes.io/docs/))
    *   Tons of work to do here. 
*   Bootcamp pages integrated into site:
    *   http://kubernetes.io/docs/tutorials/getting-started/create-cluster/

Agenda



*   Launch process for upcoming releases
    *   [Keep the spreadsheet](https://docs.google.com/spreadsheets/d/1PcAW-ZQnzlwEs59ygHdFA-72n6FUcGIc2lx4IeoXvco/edit?ts=57ced759#gid=0)
        *   Sync with [Community Spreadsheet](https://docs.google.com/spreadsheets/d/1MeTQbtSiTCoQ74zjEGRI1rXRSP0Tzmo9-kMJSK_HGoA/)
    *   Communicate out which branch to use
        *   Delete the version branch immediately after launch? (most people say yes)
        *   Number each branch
*   Pointing to the “right” version of kubernetes:
    *   [https://groups.google.com/forum/#!topic/kubernetes-sig-docs/d5vB92b4HaI](https://groups.google.com/forum/#!topic/kubernetes-sig-docs/d5vB92b4HaI)
    *   Toolset with the documentation deploys 1.3
    *   Please push to master :)
*   Issues:
    *   PR template is in!
    *   Jeff Menoza needs to merge template with “create issue” button
    *   Proposal for issues will be out for review. 

Kudos



*   Phillip Wittrock
*   Devin Donnelly

**9/20/2016**

Agenda



*   1.4 Delayed until next week ([1.4 Spreadsheet)](https://docs.google.com/spreadsheets/d/1PcAW-ZQnzlwEs59ygHdFA-72n6FUcGIc2lx4IeoXvco/edit?ts=57ced759#gid=0)
*   [Doc Templates ](https://github.com/kubernetes/kubernetes.github.io/pull/1205)
    *   Any more feedback?
    *   Still adding canonical examples
    *   Outstanding questions for those TBD by jeffmendoza
*   Bootcamp integrated into master ([pending PR](https://github.com/kubernetes/kubernetes.github.io/pull/1203))
    *   Not using the templates
    *   Keeping the existing CSS
*   Issue Template update?
    *   PR filed
        *   CLA isn’t approved. Weird.
        *   Is your github account under your company’s email?
    *   Next Steps:
        *   Submit PR
        *   AI: Develop proposal for handling backlog of issues (Abby)
        *   Run it by Frank for review
*   Dead link report:
    *   AI: Regenerate the report, mark what will fixed in the upcoming release (1.4), focus on fixing manually generated content

Updates:



*   New “edit docs” button at the bottom of the page
*   Fix in for generating API and Kubectl docs (no more broken links! Won’t get in until 1.4).
*   Kudos to Abby for creating the PR for issue templates

**9/13/2016**

Agenda



*   Update on Netlify! (Devin)
*   Dan Paik - new PM on Gcloud/K8s - will be working on cluster deployments, GKE UI, and PetSet deployments

Updates:



*   “I wish” is gone.  “File an issue” instead. (thanks Jeff!)
    *   Still need guidelines to close issues :)
*   Status of issue templates and guidelines? (Abby?)
    *   Took a stab at [making an issue template](https://docs.google.com/document/d/1aZ6vlxOhNtKNhNOUx_KTUw1bXCPPRH2_TJsNRz7CA58/edit?ts=57d83501) (in Google Doc)
    *   Feedback:
        *   Keep existing title with URL
        *   Specify feature vs. bug (new doc vs. doc update)
    *   File a PR :)
*   [1.4 Process](https://docs.google.com/spreadsheets/d/1PcAW-ZQnzlwEs59ygHdFA-72n6FUcGIc2lx4IeoXvco/edit?ts=57ced759#gid=0)
*   [Templates ](https://github.com/kubernetes/kubernetes.github.io/pull/1205)(Steve)
    *   - Netlify
*   [SEO Issues with K8s](https://docs.google.com/document/d/19A8GL7lNpvrVZJyPLHYgM7axcw6PWxT1yjQydno58L0/edit#heading=h.ucf6rfasc6ai)

**August 30, 2016**

Agenda:



*   Customer Feedback:
    *   “I wish”
        *   Make an issue / Make an edit
        *   Remove the functionality, replace with make an issue/made edits
        *   AI: Jeff Mendoza
    *   Star ratings vs. Thumbs Up
*   Issues :(
    *   300 issues
    *   Write up a md file for issue guidelines?
    *   Issue templates?
        *   AI: Abby: 
            *   Issue guidelines
            *   Issue template
    *   Close issues that don’t conform
*   Steve chat [through new templates](https://github.com/kubernetes/kubernetes.github.io/tree/docsv2/_includes/templates). 
*   1.4 Features and Publishing
    *   Alpha: [14 Open Issues](https://github.com/kubernetes/features/issues?utf8=%E2%9C%93&q=is%3Aissue%20is%3Aopen%20label%3Aalpha-in-1.4%20)
    *   Beta: [8 Open Issues](https://github.com/kubernetes/features/issues?utf8=%E2%9C%93&q=is%3Aissue%20is%3Aopen%20label%3Abeta-in-1.4%20)
    *   Stable: [5 Open Issues](https://github.com/kubernetes/features/issues?utf8=%E2%9C%93&q=is%3Aissue%20is%3Aopen%20label%3Astable-in-1.4%20)
    *   Publishing process?
        *   Current release branch 
        *   Under development release
        *   Make a rule that any change in the current version has to be made in the “under development” version.

**August 23, 2016**

Agenda



*   Introductions - New Members
    *   Jared, Abby (Apprenda), Frank Solomon (Blogger, SWE)
*   Items from the past meeting: 
    *   Third party docs?  Process for moving them?  Plan?  
        *   [Issue in current UX Backlog](https://github.com/kubernetes/kubernetes/milestone/24)
*   [Dead Link Report](https://docs.google.com/spreadsheets/d/1vLmdCSZngIRUtBcG04QGWSQ42Pr3198n-RMctq8Kv7I/edit#gid=969833937) 
    *   If you have time please fix!
    *   Please keep the status up to date
*   Kubernetes Bootcamp (Beta)
    *   [https://kubernetesbootcamp.github.io/kubernetes-bootcamp/](https://kubernetesbootcamp.github.io/kubernetes-bootcamp/)
    *   Please review if you have a chance

**August 16, 2016**

Kickoff Meeting!

Agenda



*   Introductions - New Members
    *   Jared (TW, Manager), Phillip (SWE, UX), Devin (TW), Steve (TW), Josh (SWE), Dan Romlein (UX, Apprenda), Abby (TW, Apprenda), Sarah Novotny (Community), Janet (SWE, UX), Frank, Morgan (TW, IBM)
*   Current state of the docs (Jared)
    *   John Mulhausen has left the group
    *   Devin and Steve are joining full time from Google
    *   We want to make things easier for users navigating the site, and contributors contributing to the site. 
*   Doc goals and doc plans (Devin)
    *   Building on John’s work, building new processes. 
    *   Some findability and readability improvements in the short term. 
    *   Significant problems with findability.
        *   Title != URL != Place in table of contents.  
    *   Good content with some better organization
        *   For example:
            *   Glossary is a collection of “conceptual articles”.  Should better explain the “kubernetes model”
            *   Also, a good place for new diagrams.
    *   Better guidance on audience.  Explicitly state who our audience is:
        *   App Builders - Just running an application on top of K8s. Probably don’t want/need to know too much about what’s “under the hood”
        *   Container Maintainers/Builders - Want to know the deeper aspects of K8s. 
    *   Build out the contributor 
        *   Style Guide (titling docs, content fit into categories, tone and term usage to unify content and fix SEO) 
        *   Templates - Build out 
        *   Contribution process
    *   Questions:
        *   Third party docs?  How do we build and maintain them?
            *   Proposal - that content should not live inside the core repo, since we can’t maintain it well. 
            *   No objections to this proposal so far....  Need to have a way to separate this process out and find new homes for it. 
            *   Need an owner for this process. 

Additional Agenda Items?

*   New contributor onboarding process (Phillip)
    *   **_Style guide (Josh)_**
    *   Initial tasks for new contributors
    *   Repo Issues list
        *   Github Issues that writers take on?
    *   Getting help
        *   Slack Channel (#sig-docs) for folks who need help or guidance. 
        *   Feel free to ping 
        *   To login to Zoom: Phillip directly. 
*   Diagrams - How should we create these? (Phillip)

Issues, suggestions, concerns, kudos?

*   Kudos:  Thanks to Phillip Wittrock and Kelsey Hightower for closing out our list of outstanding docs PRs