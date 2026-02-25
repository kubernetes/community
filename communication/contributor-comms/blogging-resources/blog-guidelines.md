# Contributor Comms Blog Guidelines

This initiative falls under the [Contributor Comms Charter](./CHARTER.md).

We are looking for Kubernetes-curious community members who are
**interested in writing** and **care about getting the word out** to
our huge community of users, developers, and contributors of all
types. Here's how to get involved.

## Requested Content

We are looking for content related to the contributor experience and
with increasing the visibility of Kubernetes and how it is developed:
this includes interviews with SIGs, articles on how to better use
existing tools and processes, and in general tips and suggestions on
how to collaborate.

Other types of content, like Kubernetes capabilities, tutorials, and
technical articles, are better suited for the [SIG Docs blogging
initiative](/sig-docs/blog-subproject/README.md).

## Where to publish

As mentioned, the focus of the Contributor Experience articles is
targeted at those that contribute to Kubernetes, but sometimes it's
not obvious where a specific theme will fit. The following are the
most common situations:

1. Article is just for [k8s.dev](http://k8s.dev/blog): this is when it
   is relevant for the contributor community, and not necessarily for
   Kubernetes end-users. An example is an article explaining how to
   use some specific tool or automation that helps with the Kubernetes
   development process.
2. Article is just for [kubernetes.io](https://kubernetes.io/blog/):
   when the article targets Kubernetes end-users, and not specifically
   the contributor community. Examples include most technical articles
   on Kubernetes features, updates on new features and deprecations,
   etc.
3. Article is relevant for both: sometimes, an article will be
   relevant to both the Kubernetes end-users, and the contributor
   community. Examples include interviews with SIGs and WGs, articles
   on technical aspects that are important for the contributor
   community, etc.

The decision on what is the right option will be made jointly by the
SIG Contribex Comms and the SIG Docs Blogging editorial team: as a
content writer you shouldn't be overly concerned about it, except in
how it can change the approval process, as described below.

## Submission and review process

The quickest way to get involved is to let the team in
[#sig-contribex-comms](https://kubernetes.slack.com/archives/C03KT3SUJ20)
know that you have an idea for an article; the team will identify the
best target for your submission and liaison with the necessary teams,
if needed. To reduce the amount of editing done directly in GitHub, a
two-stage approach is highly recommended.

This process is initiated in
[#sig-contribex-comms](https://kubernetes.slack.com/archives/C03KT3SUJ20)
and uses the processes from the SIG Docs blog
[subproject](/sig-docs/blog-subproject/README.md), and is broadly as
follows:

1. Present your idea to the community, by going to the
   [#sig-contribex-comms](https://kubernetes.slack.com/archives/C03KT3SUJ20)
   Slack channel, or by joining the [weekly
   meeting](https://docs.google.com/document/d/1KDoqbw2A6W7rLSbIRuOlqH8gkoOnp2IHHuV9KyJDD2c). This
   will make it easier to coordinate effort and avoid duplicate
   effort, as well as to gather initial suggestions around the article
   scope.
2. The submission idea will be reviewed by the team, including the
   decision on where to publish it; someone from the SIG Contribex
   Comms team will reach out to the
   [#sig-docs-blog](https://kubernetes.slack.com/archives/CJDHVD54J)
   editorial team to clarify if the content is adequate for
   republishing in the main Kubernetes blog. At this stage an editor
   should be assigned to follow-up the process with you.
3. Create your proposal draft in [Google
   Docs](https://docs.google.com/) or HackMD (https://hackmd.io), and
   ask for a review in [the
   channel](https://kubernetes.slack.com/archives/C03KT3SUJ20). This
   will facilitate easier editing, especially if major changes or
   restructuring is needed. Take into account the [documentation style
   guide](https://kubernetes.io/docs/contribute/style/style-guide/):
   these guidelines can help in improving the readability of your
   article, especially in terms of the use of Kubernetes terminology.
   You can also read
   [Submitting blog posts and case studies](https://kubernetes.io/docs/contribute/new-content/blogs-case-studies/)
   for extra context.
4. Once you have reflected any feedback in the proposal draft,
   announce that the article is ready for submission (again, in the
   channel or in one of the weekly meetings): the assigned editor will
   use the final text to open the PR, adding you as a Co-author.
5. (Optional) If the submission will be mirrored in the main
   Kubernetes site, a second PR will be opened by the editor, but on
   the main repository. The content of both should be the same, with
   the applicable differences (e.g., file location, metadata). The
   review process primarily happens in the `contributor-site` PR, with
   all the changes then copied over to the `website` after approval.

For now, our official process is to use [SIG Docs'
system](/sig-docs/blog-subproject/README.md), with one change: instead
of directly creating the file in the Kubernetes site repository, as
instructed above it's initially created in the
[contributor-site](https://github.com/kubernetes/contributor-site), in
the appropriate folder (i.e. the right year in
`contributor-site/content/en/blog/`).

This will lead to an initial review process before it gets mirrored to
the main Kubernetes site.

### Editor instructions

Once the text is final, an editor will open the PR. This facilitates
the approval process and prevents articles with massive restructuring
or changes needed to be submitted to GitHub, something that makes the
review process substantially more difficult.

In order to keep the authorship information (which will make the
submission count towards the contribution of the article author),
editors must [add the original author as a
co-author](https://docs.github.com/en/pull-requests/committing-changes-to-your-project/creating-and-editing-commits/creating-a-commit-with-multiple-authors). This
is done by adding `Co-authored-by: original-author-name
<original-author@example.com>` to the commit message.

The number of PRs you open depends on where the article will be published:

1. If it's solely for the Contributor site: the PR should be opened in
   the
   [contributor-site](https://github.com/kubernetes/contributor-site),
   after which the process ends.
2. If it's to be mirrored in the main Kubernetes blog: after the
   previous step, a new PR is opened on
   [kubernetes/website](https://github.com/kubernetes/website),
   mentioning the original PR. Reviewers from SIG Docs Blog will,
   in this case, already been notified and involved in the initial PR.


## Blogger Expectations, Responsibilities, and Information

Anyone is welcome to contribute when they have time.

If you would like to be listed as a member of the team, here are the expectations:

1. Be prepared to write one blog a quarter and participate in edits to
   other articles. The time commitment is typically 5-10 hours per
   quarter depending on the number of blog posts in the review queue.
2. Bloggers are expected to attend at least one Contributor Comms team
   meeting a month or check-in to remain active.
3. Remain non-partial: if you receive a request to write about a
   project, an individual, or a group of people from your employer,
   you should ask an impartial blogger to write it.
4. As with all contribution to Kubernetes, adhere to the [code of
   conduct](/code-of-conduct.md), values, and principles of the
   project.

## How to Write an Effective Blog

Keep the following points in mind as you write in order to speed up the review process:

* Use inclusive language understandable by everyone
  * Rephrase gendered pronouns (change "he" or "she" to "they" or
    adjust to remove)
  * Remember nothing is simple when you're starting out (remove
    "just," "simply", and "easy")
  * Define terminology or acronyms (do not assume people know what a term means)
  * Shy away from jargon and colloquial expressions
  * Write clearly and avoid ambiguous sentences
* Emphasize the things you want readers to remember; tell a story
  * Stay on topic and stick to the facts
  * Design a beginning, middle, and end to your story with a clear call to action
  * Provide evidence and data where applicable, to back up your message
* Make the article visually appealing
  * Include at least one image (and use public domain or Creative
    Commons licensed ones)
  * Prefer inclusive images like those from
    [WOCinTech](https://www.flickr.com/photos/wocintechchat/) and
    [Queer in
    Tech](https://www.flickr.com/photos/mapbox/albums/72157713100349311)
  * Find images on sites like [Creative
    Commons](https://search.creativecommons.org/),
    [Pexels](https://www.pexels.com/public-domain-images/), and
    [Unsplash](https://unsplash.com/images/stock/public-domain))
* Be accountable and honest as an author
  * Remove anything that lacks adequate evidence
  * Avoid interjecting personal reactions
  * Ensure that the blog post is reviewed by anyone being mentioned in the piece
  * As the author, never talk about your employer, sell, promote, or
    pitch; this is about upstream community endeavours and the
    individuals and groups that create it
* Follow the [documentation style guide](https://kubernetes.io/docs/contribute/style/style-guide/);
  for blog articles these are informative guidelines rather than anything more strict, but
  it's still good to follow them where appropriate.

## Further Recommendations

The following are helpful resources for authoring articles:

* [Creating Quality Content (For Search Engines and
  People)](https://moz.com/blog/quality-blog-content)
* [How to write effective documentation for your open source
  project](https://opensource.com/article/20/3/documentation)
