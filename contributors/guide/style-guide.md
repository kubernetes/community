---
title: "Documentation Style Guide"
weight: 1
slug: "documentation-style-guide" 
---

This style guide is for content in the Kubernetes github [community repository].
It is an extension of the [Kubernetes documentation style-guide].

These are **guidelines**, not rules. Use your best judgement.

- [Cheatsheet](#cheatsheet)
- [Content design, formatting, and language](#content-formatting-and-language)
  - [Contact information](#contact-information)
  - [Dates and times](#dates-and-times)
  - [Diagrams, images and other assets](#diagrams-images-and-other-assets)
  - [Document Layout](#document-layout)
  - [Formatting text](#formatting-text)
  - [Language, grammar, and tone](#language-grammar-and-tone)
  - [Moving a document](#moving-a-document)
  - [Punctuation](#punctuation)
  - [Quotation](#quotation)
- [Markdown formatting](#markdown-and-formatting)
  - [Code Blocks](#code-blocks)
  - [Emphasis](#emphasis)
  - [Headings](#headings)
  - [Horizontal Lines](#horizontal-lines)
  - [Line Length](#line-length)
  - [Links](#links)
  - [Lists](#lists)
  - [Metadata](#metadata)
  - [Tables](#tables)
- [Attribution](#attribution)


## Cheatsheet

### Cheatsheet: Content Design, Formatting, and Language

**[Contact Information:](#contact-information)**
- Use official Kubernetes contact information.

**[Dates and Times:](#dates-and-times)**
- Format dates as `month day, year`. (December 13, 2018)
- When conveying a date in numerical form, use [ISO 8601] Format: `yyyy-mm-dd`.
- Use the 24 hour clock when referencing time.
- Times for single events (example: KubeCon) should be expressed in an absolute
  time zone such as Pacific Standard Time (PST) or Coordinated Universal Time
  (UTC).
- Times for reoccurring events should be expressed in a time zone that follows
  Daylight Savings Time (DST) such as Pacific Time (PT) or Eastern Time (ET).
- Supply a link to a globally available time zone converter service.
  - `http://www.thetimezoneconverter.com/?t=<TIME REFERENCE>&tz=<TZ REFERENCE>`

**[Diagrams, Images and Other Assets:](#diagrams-images-and-other-assets)**
- Images and other assets should be stored in the same directory as the document
  that is referencing it.
- Filenames should be lowercase and descriptive of what they are referencing.
- Avoid excessively large images or include a smaller one while linking to a
  higher resolution version of the same image.
- Use the [Kubernetes icon set] for architectural diagrams.

**[Document Layout:](#document-layout)**
- Documents should follow the general template of:
  - Document metadata (if appropriate).
  - Title in `H1` (a single `#`).
  - A brief description or summary of the document.
  - A table of contents.
  - The general body of document.
- Do not repeat content. Instead link back to the canonical source.
- Large content or topic shifts should be separated with a horizontal rule.

**[Formatting Text:](#formatting-text)**
- API objects:
  - Follow the established [API naming convention] when referring to API Objects.
  - Do not split API object names into their components.
  - Use `code` style for API objects or object parameters.
- Use **bold text** for user interface elements.
- Use _italics_ to emphasize a new topic or subject for the first time.
- Use angle brackets (`<` and `>`) to enclose a placeholder reference.
- Apply `code` styling to:
  - Filenames, directories, and paths.
  - Command line examples and flags.
  - Object field names.

**[Language, Grammar and Tone:](#language)**
- Documentation should be written in English.
- Prefer an active voice and present tense when possible.
- Use simple and direct language.
- Use gender-neutral language.
- Avoid personal pronouns ("I," "we," "us," "our," and "ours").
- Address the reader as "you" instead of "we".
- Do not use Latin phrases.
- Avoid jargon and idioms.
- If using acronyms, ensure they are clearly defined in the same document.
- If using an abbreviation, spell it out the first time it is used in the
  document unless it is commonly known. (example: TCP/IP)

**[Moving a Document:](#moving-a-document)**
- Use `[git-mv]` to move documents.
- Commit moved documents separately from any other changes.
- When a document has moved, leave a tombstone file with a removal date in its
  place.

**[Punctuation:](#punctuation)**
- Do not use punctuation in headings.
- End full sentences with a period.
  - **Exception:** When a sentence ends with a URL or if the text would be
    unclear if the period is a part of the previous object or word.
- Add a single space after a period when beginning a new sentence.
- Avoid usage of exclamation points unless they are a part of a code example.
- Use an [Oxford comma] when a list contains 3 or more elements.

**[Quotation:](#quotes)**
- Use double-quotation marks (`" "`) over single-quotation marks (`' '`).
  - **Exception:** In code snippets where quotation marks have specific meaning.
  - **Exception:** When nesting quotation marks inside another set of quotation
    marks.
- Punctuation should be outside of quotation marks following the international
  (British) standard.


### Cheatsheet: Markdown

**[Code Blocks:](#code-blocks)**
- When possible, reference the language at the beginning of a Code Block.
- When a code block is used to reference a shell, do not include the command
  prompt (`$`).
  - **Exception:** When a code block is used to display raw shell output.
- Separate commands from output.

**[Emphasis:](#emphasis)**
- Use two asterisks (`**`) for **Bold** text.
- Use an underscore (`_`) for _Italics_.
- Use two tildes (`~~`) for ~~Strikethrough~~.

**[Headings:](#headings)**
- Use a single `H1` (`#`) Heading per document.
  - **Exception:** `H1` may be used multiple times in the same document when
    there is a large content shift or "chapter" change.
- Follow the Header hierarchy of `H2` > `H3` > `H4` > `H5` > `H6`.
- Use title-case capitalization.
  - Capitalize the first word.
  - Capitalize all nouns, verbs, adjectives, adverbs, and pronouns.
  - Capitalize all words of four letters or more.
  - Use lowercase words of three or fewer letters.
- Avoid using special characters.
- Leave exactly 1 new line after a heading.
- Avoid using links in headings.

**[Horizontal Rules:](#horizontal-lines)**
- Use three dashes (`---`) to denote a horizontal rule.
- Use a horizontal rule (`---`) to logically separate large sections.

**[Line Length:](#line-length)**
- Prefer an 80 character line limit.

**[Links:](#links)**
- Prefer using reference style links over inline style links.
- When linking within the same directory, use a relative link.
- When linking to a document outside of the current directory, use the absolute
  path from the root of the repository.
- When linking to a file in another Kubernetes github repository, use the
  `k8s.io` url shortener.
  - git.k8s.io -> github.com/kubernetes
  - sigs.k8s.io -> github.com/kubernetes-sigs

**[Lists:](#lists)**
- Capitalize the first character of each entry unless the item is explicitly
  case sensitive.
- End each entry with a period if it is a sentence or phrase.
- Use a colon (`:`) to separate a list item name from the explanatory text.
- Leave a blank line after each list.
- Use `-` for unordered lists.
- For ordered lists repeating `1.` may be used.
- When inserting a code block into an ordered list, indent (space) an additional
  two times.

**[Metadata:](#metadata)**
- If the document is intended to be surfaced on the Contributor Site; include a
  yaml metadata header at the beginning of the document.
- Metadata must include the `title` attribute.

**[Tables:](#tables)**
- Use tables for structured information.
- Tables do not need to adhere to the suggested line length.
- Avoid long inline links.
- Do not use excessively wide tables.

---

## Content Design, Formatting, and Language

### Contact Information

- Use official Kubernetes contact information.
  - Use official community contact email addresses. There should be no personal
    or work contact information included in public documentation; instead use
    addresses like the [SIG Google groups] or managed accounts such as
    community@kubernetes.io.
  - **Good example:** community@kubernetes.io
  - **Bad example:** bob@example.com


### Dates and Times

The Kubernetes Contributor Community spans many regions and time zones.
Following a consistent pattern and avoiding shorthand improves the readability
for every member.

- Format dates as `month day, year`. (December 13, 2018)
  - **Good example:** October 24, 2018
  - **Bad example:** 10/24/18
- When conveying a date in numerical form, use [ISO 8601] Format: `yyyy-mm-dd`.
  - **Good example:** 2018-10-24
  - **Bad example:** 10/24/18
- Use the 24 hour clock when referencing time.
  - **Good example:** 15:30
  - **Bad example:** 3:30pm
- Times for single events (example: KubeCon) should be expressed in an absolute
  time zone such as Pacific Standard Time (PST) or Coordinated Universal Time
  (UTC).
  - **Good example:** The Seattle Contributor Summit starts at 9:00 PST
  - **Bad example:** The Seattle Contributor Summit starts at 9:00 PT
- Times for reoccurring events should be expressed in a time zone that follows
  Daylight Savings Time (DST) such as Pacific Time (PT) or Eastern Time (ET).
  - Times that follow DST are used as they adjusts automatically. If UTC or
    other non-DST compatible time zones were used, content would have to be
    updated multiple times per year to adjust times.
  - **Good example:** 13:30 PT
  - **Bad example:** 16:30 EST
- Supply a link to a globally available time zone converter service.
  - `http://www.thetimezoneconverter.com/?t=<TIME REFERENCE>&tz=<TZ REFERENCE>`

  ```
  The weekly SIG meeting is at [13:30 PT].

  [13:30 PT]: http://www.thetimezoneconverter.com/?t=13:30&tz=PT%20%28Pacific%20Time%29
  ```


### Diagrams, Images and Other Assets

- Images and other assets should be stored in the same directory as the document
  that is referencing it.
- Filenames should be lowercase and descriptive of what they are referencing.
  - **Good example:** `deployment-workflow.jpg`
  - **Bad example:** `image1.jpg`
- Avoid excessively large images or include a smaller one while linking to a
  higher resolution version of the same image.
- Use the [Kubernetes icon set] for architectural diagrams.


### Document Layout

Adhering to a standard document layout ensures that each page can intuitively
be navigated once a reader is familiar with the standard layout.

- Documents should follow the general template of:
  - Document metadata (if appropriate).
  - Title in `H1` (a single `#`).
  - A brief description or summary of the document.
  - A table of contents.
  - The general body of document.
- Do not repeat content. Instead link back to the canonical source.
  - It is easy for content to become out of sync if it is maintained in
    multiple places. Linking back to the canonical source ensures that the
    documentation will be accurate and up to date.
- Large content or topic shifts should be separated with a horizontal rule.


### Formatting Text

The formatting guidelines have been selected to mirror or augment the
[Kubernetes documentation style-guide]. Remaining consistent across the
different content sources improves the overall readability and understanding of
the documentation being presented in addition to giving the project a unified
external appearance.

- API objects:
  - Follow the established [API naming convention] when referring to API Objects.
  - Do not split API object names into their components.
    - **Good example:** A `Pod` contains a `PodTemplateSpec`.
    - **Bad example:** A `Pod` contains a `Pod Template Spec`.
  - Use `code` style for API objects or object parameters.
    - **Good example:** A `Deployment` contains a `DeploymentSpec`.
    - **Bad example:**  A Deployment contains a DeploymentSpec.
- Use angle brackets (`<` and `>`) to surround a placeholder references.
  - **Good example:** `kubectl describe pod <pod-name>`
  - **Bad example:** `kubectl describe pod pod-name`
- Use **bold text** for user interface elements.
  - **Good example:** Select **Other**.
  - **Bad example:** Select "Other".
- Use _italic text_ to emphasize a new subject for the first time.
  - **Good example:** A _cluster_ is a set of nodes.
  - **Bad example:** A "cluster" is a set of nodes.
- `Code` styling should be applied to:
  - Filenames, directories, and paths.
    - **Good example:** The default manifest path is `/etc/kubernetes/manifests`.
    - **Bad example:** The default manifest path is /etc/kubernetes/manifests.
  - Command line examples and flags.
    - **Good example:** The flag `--advertise-address` is used to denote the
      IP address on which to advertise the apiserver to members of the cluster.
    - **Bad example:** The flag --advertise-address is used to denote the IP
      address on which to advertise the apiserver to members of the cluster.
  - Object field names.
    - **Good example:** Set the `externalTrafficPolicy` to Local.
    - **Bad example:** Set the externalTrafficPolicy to Local.


### Language, Grammar and Tone

- Documentation should be written in English.
- Prefer an active voice and present tense when possible.
  - Active voice is when the subject of the sentence performs the action.
    Whereas with passive voice the subject receives the action. Writing with an
    active voice in mind easily conveys to the reader who or what is performing
    the action.
  - **Good example:** Updating the Deployment triggers a new ReplicaSet to be
    created.
  - **Bad example:** A ReplicaSet is created by updating the Deployment.
- Use simple and direct language.
  - Avoid using unnecessary or extra language. Be straightforward and direct.
  - **Good example:** Wait for the Pod to start.
  - **Bad example:** Please be patient and wait for the Pod to start.
- Use gender-neutral language.
  - Avoid gendered pronouns preferring the [singular "they"][singular-they]
    unless referring to the person's by their preferred gender. For further
    information on the subject, see [Microsoft's guide to bias-free communication]
    and [Wikipedia's entry for the Singular they].
  - **Good example:** chair or moderator
  - **Bad example:** chairman
- Avoid personal pronouns ("I," "we," "us," "our," and "ours").
  - In most cases personal pronouns should be avoided as they can lead to
    confusion regarding who they are referring to.
  - **Good example:** The release-team shepherded the successful release of 1.13.
  - **Bad example:** We shepherded the successful release of 1.13.
- Address the reader as "you" instead of "we".
  - Addressing the reader directly using "you" clearly denotes the target.
    There is no confusion as there would be with "we" or "us".
  - **Good example:** You will create a new cluster with kubeadm.
  - **Bad example:** We will create a new cluster with kubeadm.
- Do not use Latin phrases.
  - [Latin phrases] can make it difficult for readers not familiar with them to
    grasp their meaning.
  - Some useful alternatives include:

    | Latin Phrase | Alternative |
    |:------------:|:-----------:|
    |     e.g.     | for example |
    |    et al.    |  and others |
    |     i.e.     |   that is   |
    |     via      |    using    |

  - **Good example:** For example Deployments, ReplicaSets...
  - **Bad example:** e.g. Deployments, ReplicaSets...
- Avoid jargons and idioms.
  - Jargon and idioms tend to rely on regional or tribal knowledge. They make
    it difficult to understand for both newcomers and those whose native
    language is something other than English. They should be avoided when
    possible.
  - **Good example:** Internally, the kube-apiserver...
  - **Bad example:** Under the hood the kube-apiserver...
  - **Good example:** We will start the project in early 2019.
  - **Bad example:** We will kick off the initiative in 2019.
- If using an abbreviation, spell it out the first time it is used in the
  document unless it is commonly known. (example: TCP/IP)
  - Abbreviations in this context applies to abbreviations, acronyms and
    initialisms.
  - **Good example:** A _CustomResourceDefinition_ (CRD) extends the Kubernetes
    API.
  - **Bad example:** A CRD extends the Kubernetes API.


### Moving a Document

- Use `[git-mv]` to move documents.
  - `git-mv` will safely move/rename a file, directory, or symlink and
    automatically update the git index.
  - **Good example:** `git mv /old/mydoc.md /new/mydoc.md`
  - **Bad example:** `mv /old/mydoc.md /new/mydoc.md`
- Commit moved documents separately from any other changes.
  - A separate commit clearly preserves the history of the relocated documents
    and makes it easier to review.
- When a document has moved, leave a tombstone file with a removal date in its
  place.
  - Tombstones function as a pointer and give users a time to update their own
    documentation and bookmarks. Their usefulness is time-bounded and should be
    removed when they would logically no longer serve their purpose.
    ```markdown
    This file has moved to https://git.k8s.io/community/contributors/guide/README.md.

    This file is a placeholder to preserve links.  Please remove after 2019-03-10 or the release of kubernetes 1.10, whichever comes first.
    ```


### Punctuation

- Do not use punctuation in headings.
- End full sentences with a period.
  - **Exception:** When a sentence ends with a URL or if the text would be
    unclear if the period is a part of the previous object or word.
- Add a single space after a period when beginning a new sentence.
- Avoid usage of exclamation points unless they are a part of a code example.
- Use an [Oxford comma] when a list contains 3 or more elements.
  - **Good example:** Deployments, ReplicaSets, and DaemonSets.
  - **Bad example:** Deployments, ReplicaSets and DaemonSets.


### Quotation

- Use double-quotation marks (`" "`) over single-quotation marks (`' '`).
  - **Exception:** In code snippets where quotation marks have specific meaning.
  - **Exception:** When nesting quotation marks inside another set of quotation
    marks.
- Punctuation should be outside of quotation marks following the international
  (British) standard.


---


## Markdown Formatting

### Code Blocks

- When possible, reference the language at the beginning of a Code Block.
  - The two markdown renderers used by the Kubernetes community
    ([GitHub][gh-code-hl-list] and [Hugo][hugo-code-hl-list]) support code
    highlighting. This can be enabled by supplying the name of the language
    after the three back-ticks (`` ``` ``) at the start of a code block.
  - **Good example:**
    `````
    ```go
    import (
      "fmt"
      ...
    )
    ```
    `````
  - **Bad example:**
    `````
    ```
    import (
      "fmt"
      ...
    )
    ```
    `````
- When a code block is used to reference a shell, do not include the command
  prompt (`$`)
  - When a code block is referencing a shell, it is implied that it is a
    command prompt. The exception to this is when a code block is being used
    for raw shell output such as debug logs.
  - **Good example:**
    ```
    kubectl get pods -o wide
    ```
  - **Bad example:**
    ```
    $ kubectl get pods -o wide
    ```
- Separate commands from output.
  - Separating the command from the output makes both the command and output
    more generally readable.
  - **Good example:**
    ```
    kubectl get pods
    ```
    ```
    NAME     READY     STATUS    RESTARTS   AGE    IP           NODE
    nginx    1/1       Running   0          13s    10.200.0.4   worker0
    ```
  - **Bad example:**
    ```
    kubectl get pods
    NAME     READY     STATUS    RESTARTS   AGE    IP           NODE
    nginx    1/1       Running   0          13s    10.200.0.4   worker0
    ```


### Emphasis

Markdown has multiple ways of indicating each type of emphasis. Adhering to a
standard across documentation improves supportability.

- Use two asterisks (`**`) for **Bold** text.
  - **Good example:** `This is **bold** text.`
  - **Bad example:** `This should not be used for __bold__.`
- Use an underscore (`_`) for _Italics_.
  - **Good example:** This is _italics_.`
  - **Bad example:**  This should not be used for *italics*.`
- Use two tildes (`~~`) for ~~Strikethrough~~.
  - **Good example:** `This is ~~strikethrough~~`
  - **Bad example:** `This should not be used for ~strikethrough~.`


### Headings

 Adhering to a standard across documentation improves both readability and
 overall supportability across multiple documents.

- Use a single `H1` (`#`) Heading per document.
  - **Exception:** `H1` may be used multiple times in the same document when
    there is a large content shift or "chapter" change.
- Follow the Header hierarchy of `H2` > `H3` > `H4` > `H5` > `H6`.
- Use title-case capitalization.
  - Capitalize the first word.
  - Capitalize all nouns, verbs, adjectives, adverbs, and pronouns.
  - Capitalize all words of four letters or more.
- Avoid using special characters.
- Leave exactly 1 new line after a heading.
- Avoid using links in headings.


### Horizontal Rules

Markdown has multiple ways of indicating a horizontal rule. Adhering to a
standard across documentation improves supportability.

- Use three dashes (`---`) to denote a horizontal rule.
  - **Good example:** `---`
  - **Bad example:** `===`
- Use a horizontal rule (`---`) to logically separate large sections.


### Line Length

- Prefer an 80 character line limit.
  - There is no specific general best practice for Markdown line length. The
    commonly used 80 character guideline is preferable for general text review
    and editing.


### Links

Markdown provides two primary methods to link to content: inline links and
relative links. However, how and what they're being linked to can vary widely.

- Prefer using reference style links over inline style links.
  - Reference links are shorter and easier to read. They have the added benefit
    of being reusable throughout the entire document.
  - The link itself should be at the bottom of the document. If the document is
    large or covers many topics, place the link at the end of the logical
    chapter or section.
  - **Example:**
    ```
    See the [Code of Conduct] for more information.

    [code of conduct]: https://git.k8s.io/community/code-of-conduct.md
    ```
  - **Example:**
    ```
    See the [Code of Conduct][coc] for more information.

    [coc]: https://git.k8s.io/community/code-of-conduct.md
    ```
- When linking within the same directory, use a relative link.
  - Links to files within the same directory are short and readable already.
    They do not warrant expanding the full path.
  - When the file is referenced multiple times within the same document,
    consider using a reference link for a quicker shorthand reference.
  - **Example:**
    ```
    See the [Code of Conduct](code-of-conduct.md) for more information
    ```
  - **Example:**
    ```
    See the [Code of Conduct][coc] for more information

    [coc]: code-of-conduct.md
    ```
- When linking to a document outside of the current directory, use the absolute
  path from the root of the repository.
  - Using the absolute path ensures that if the source document is relocated,
    the link to the target or destination document will remain intact and not
    have to be updated.
  - **Example:**
    ```
    See the [Coding Convention] doc for more information.

    [Coding Convention]: /contributors/guide/coding-conventions.md
    ```
- When linking to a file in another Kubernetes github repository, use the
  `k8s.io` url shortener.
  - The shorthand version will auto-expand linking to documents within the
    master branch and can be used for multiple purposes.

    |      Short URL      |            Expanded URL            |
    |:-------------------:|:----------------------------------:|
    |  https://git.k8s.io |    https://github.com/kubernetes   |
    | https://sigs.k8s.io | https://github.com/kubernetes-sigs |

  - **Example:**
    ```
    The super cool [prow tool] resides in the test-infra repo under the kubernetes organization

    [prow tool]: https://git.k8s.io/test-infra/prow/README.md
    ```


### Lists

 Adhering to a standard across documentation improves both readability and
 overall supportability across multiple documents.

- Capitalize the first character of each entry unless the item is explicitly
  case sensitive.
- End each entry with a period if it is a sentence or phrase.
- Use a colon (`:`) to separate a list item name from the explanatory text.
- Leave a blank line after each list.
- Use `-` for unordered lists.
- For ordered lists a repeating `1.` may be used.
- When inserting a code block into an ordered list, indent (space) an additional
  two times.


### Metadata

- If the document is intended to be surfaced on the Contributor Site; include a
  yaml metadata header at the beginning of the document.
  - If the document is to be added to the Contributor Site, adding metadata
    at the beginning of the document will improve the overall presentation of
    the information. This metadata is similar to the metadata used in the
    KEP process and is often referred to as _Frontmatter_ in common static
    site generators such as [Jekyll] and [Hugo].
  - The metadata header is a yaml block between two sets of `---`.
  - **Example:**
    ```
    ---
    title: Super Awesome Doc
    ---
    ```
- Metadata must include the `title` attribute.
  - `title` will be used as the title of the document when rendered with
    [Hugo].


### Tables

- Use tables for structured information.
  - **Example:**
    ```
    |    Column 1    |    Column 2    |    Column 3    |
    |:--------------:|:--------------:|:--------------:|
    |     test 1     |     test 2     |     test 3     |
    | another test 1 | another test 2 | another test 3 |
    ```
- Tables do not need to adhere to the suggested line length.
  - Markdown tables have an inherently longer line length, and cannot be
    line wrapped.
- Avoid long inline links.
  - Long inline links can make it difficult to work with markdown tables.
    Prefer to use reference style links instead.
- Do not use excessively wide tables.
  - Large wide tables do not render well. Try to break the information down
    into something more easily presentable.


## Attribution

This style guide is heavily influenced by the great work from the content
management teams from [SIG-Docs], [Gitlab], [Google], and [Microsoft]. Without
their previous efforts this guide would not be nearly as concise as it should.

[community repository]: https://git.k8s.io/community
[Kubernetes documentation style-guide]: https://kubernetes.io/docs/contribute/style/style-guide/
[SIG Google groups]: /sig-list.md
[ISO 8601]: https://en.wikipedia.org/wiki/ISO_8601
[kubernetes icon set]: /icons
[API naming convention]: /contributors/devel/sig-architecture/api-conventions.md#naming-conventions
[singular-they]: https://en.wikipedia.org/wiki/Singular_they
[Microsoft's guide to bias-free communication]: https://docs.microsoft.com/en-us/style-guide/bias-free-communication
[Wikipedia's entry for the Singular they]: https://en.wikipedia.org/wiki/Singular_they
[Latin phrases]: https://en.wikipedia.org/wiki/List_of_Latin_abbreviations
[Oxford comma]: https://www.grammarly.com/blog/what-is-the-oxford-comma-and-why-do-people-care-so-much-about-it/
[gh-code-hl-list]: https://github.com/github/linguist/blob/master/lib/linguist/languages.yml
[hugo-code-hl-list]: http://www.rubycoloredglasses.com/2013/04/languages-supported-by-github-flavored-markdown/
[git-mv]: https://git-scm.com/docs/git-mv
[jekyll]: https://jekyllrb.com/
[hugo]: https://gohugo.io/
[gitlab]: https://docs.gitlab.com/ee/development/documentation/styleguide.html
[google]: https://developers.google.com/style/
[microsoft]: https://docs.microsoft.com/en-us/style-guide/welcome/
[sig-docs]: https://kubernetes.io/docs/contribute/style/style-guide/
