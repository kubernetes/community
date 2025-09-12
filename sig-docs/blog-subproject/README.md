# Kubernetes Blog Subproject

The Kubernetes Blog Subproject is owned by [SIG Docs](https://github.com/kubernetes/community/tree/master/sig-docs).

This section covers documentation, processes, and roles for the [Kubernetes blog](https://kubernetes.io/blog/).

## Meetings

See [meetings](https://github.com/kubernetes/community/tree/master/sig-docs#meetings) for SIG Docs


## Subproject contributors

<!-- GitHub username alphabetical order -->
- **Blog approvers:** [Bob Killen](https://github.com/mrbobbytables), [Taylor Dolezal](https://github.com/onlydole),
  [Nate Waddington](https://github.com/nate-double-u), [Tim Bannister](https://github.com/sftim)

- **Blog shadow approvers:** _no contributors_

- **Blog editors:** [Gaurav Padam](https://github.com/Gauravpadam)

- **Blog shadow editors:** _no contributors_

✨ Could **you** join the blog editorial team?

## Contact

- Slack: [#sig-docs-blog](https://kubernetes.slack.com/messages/CJDHVD54J)
- Mailing List: [blog@kubernetes.io](mailto:blog@kubernetes.io)
- Open Community Issues/PRs: [`is:open repo:kubernetes/website label:area/blog`](https://github.com/issues?q=is%3Aopen+label%3Aarea%2Fblog+repo%3Akubernetes%2Fwebsite)

## Submit a Post

Anyone can write a blog post and submit it for review. Blog posts should not be commercial in nature and should consist of content that will apply broadly to the Kubernetes community.

To propose a blog post, read [Submitting blog posts and case studies](https://k8s.io/docs/contribute/new-content/blogs-case-studies/).

### Article guidelines

**Original content only**. You cannot submit a blog article that has been published elsewhere. The Kubernetes project
makes exception to this only for articles posted to the CNCF blog or to the [Kubernetes contributor blog](https://k8s.dev/blog/).

Requested content:

- New Kubernetes capabilities
- Kubernetes projects updates
- Updates from Special Interest Groups
- Tutorials and walkthroughs
- Thought leadership around Kubernetes
- Kubernetes Partner OSS integration

Unsuitable content:

- Vendor product pitches
- Partner updates without an integration and customer story
- Syndicated posts (it's OK to localized existing articles from English)

## Review process

Once a blog post is submitted either via the form or a PR, it will be routed to the editorial team for review either via email for Google Docs or auto-assigning for a PR.

Each blog post requires a LGTM from a blog editor (or approver) and an approval by a blog approver. Blog editors will usually also get a technical review from the appropriate SIG.

_If a blog post does not contain any technical content (for example, [How You Can Help Localize Kubernetes Docs](https://kubernetes.io/blog/2019/04/26/how-you-can-help-localize-kubernetes-docs/)), the technical review can be omitted._

Articles should merge _before_ their publication date; automation picks up scheduled posts and publishes them automatically.


### Release communications

SIG Release lead on blog articles to announce Kubernetes releases, and the post-release series of articles. SIG Docs and the blog subproject support
that process and provide approvals for upcoming articles.

### Embargoed content

The blog repository on GitHub is public, therefore any content that needs to remain confidential until a certain time (for example: release posts, security vulnerabilities) should be proposed by email message to [blog@kubernetes.io](mailto:blog@kubernetes.io). If you need to, you can also send a Slack direct message to the set of blog approvers; please do this sparingly.

In your message, please note the time that the embargo will be lifted.

### SLA

Blog posts can take up to **4 weeks** to review. If you’d like to request an expedited review, please get in touch via [#sig-docs-blog](https://app.slack.com/client/T09NY5SBT/CJDHVD54J) on the Kubernetes Slack workspace.

## Ongoing blog maintenance

SIG Docs approvers for English content can approve edits after the fact such as: broken links, copy edits, etc. However, approval and editorial review for new blog posts is limited to the Blog Team.

We typically do not make edits to blog posts more than 1 years old; there is an exception for articles marked `evergreen: true` in their
[front matter](https://gohugo.io/content-management/front-matter/).

## Editorial team selection

Bloggers and reviewer responsibilities include staffing the team, with this order of fall-through in mind:

- training and selecting a successor from the current pool of role shadows
- training and selecting a successor from non-Editorial Team members
- staffing the role themselves

Ultimately, if none of these can be satisfied, responsibility falls to the SIG Docs leadership to staff the roles.

### Shadows

We are always open to adding new shadows to the editorial team roles. If you are interested in shadowing one of the roles on the team, please say Hi in
[#sig-docs-blog](https://app.slack.com/client/T09NY5SBT/CJDHVD54J) on the Kubernetes Slack workspace. Visit https://slack.k8s.io/ for an invitation if you are not already part of the workspace.

### Removing a team member

If all members of a group (eg approvers) can no longer fulfil their duties, and there is a shadow for that role, that role defaults to the shadow.
