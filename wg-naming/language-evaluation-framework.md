
# Language harm evaluation framework

## About

The language evaluation framework is a guidance document developed by the Kubernetes Naming Working Group. It outlines a structured framework for evaluating language and terminology for harm to the community. This enables the community to navigate divisive conversations with a measure of clarity.

While the document was created for an open source technology project, we feel the principles outlines are applicable to other fields as well.

## The framework

### Using the framework

The framework is divided into 3 sections: first-, second-, and third-order concerns, ranked in order of potential for harm the language could cause to the community. 

In general, first-order concerns are egregious, overt, and clearly problematic. Second-order concerns are things which indicate problematic language, but are less clearly provable. Third-order concerns indicate language that could use improvement.

For each term under evaluation, answer all questions.


When complete, look at the questions answered in the affirmative: in general, the more questions answered “yes” or “possibly”, the more likely it is that the language in question needs to be erased. If any first-order concerns are a “yes”, replace the language in question. If a significant number of second- or third- order concerns are a “yes”, strongly consider replacing the language in question. 

This framework is intentionally non-prescriptive. Bear in mind that your intention in this work is to reduce harm for the community first and foremost, and let that guide your decisions.

### First-order concerns

First-order concerns are characterized by: 

- Overtness: regardless of its use in the context of code or technology, there is little to no ambiguity outside of technology as to whether the language in question indicates harm.
- Identity-specificity: language in question specifically unambiguously identifies a group of people.

#### Is the term overtly racist?

Examples include “master/slave”.

#### Is the term overtly sexist, transphobic, or pejorative about a gender identity?

Examples do _not_ include “transclusion” of dependencies, or “binary” operators. 

#### Is the term overtly ableist, or pejorative to neurodiverse or disabled people

Examples include performing “sanity checks”. 

#### Is the term overtly homophobic? 

Examples do not include “homogenizing” or “homogenous” data. 

### Second-order concerns

Second-order concerns are characterized by: 

- Ambiguity: outside the context of code or technology, language might have connotations related to harmful scenarios like war, militarization, or policing, but the actual etymology of the term is not related to harm of a specific identity
- Lack of identity-specificity: concerns in this category do not target specific identities, or do so in a non-overt way.

#### Is the term violent?

Examples include “KILL” commands in Unix systems.

#### Is the term militaristic?

Examples include “marshal/unmarshal”.


### Third-order concerns

Third-order concerns are characterized by:

- Language focus: is the language in use a metaphor that could be described more precisely using different words?


#### Is the term evocative instead of descriptive?

Examples include “PetSet” (evocative) versus “StatefulSet” (descriptive). 

#### Is the term ambiguous?

Examples include the use of ABORT/STOP/KILL in Unix-like systems, where they map to specific behaviors, versus general usage in programming languages, where they map to different behaviors or are used interchangeably. 

## Footnotes 

### The element of time

In general, strong democratic societies become more progressive and accepting as time passes. This is a feature, not a bug. 

The result of this, for your work, is that terms that were once deemed acceptable may, at some future point, be deemed unacceptable. 

We recommend:

- Placing a date at the top of any documents/recommendations related to naming, language inclusivity, or harm reduction 
- Expecting that some of your work will need re-evaluation at a later date
- An openness to reversing decisions 


### Dealing with trolls

In the handful of months since this work began, both Kubernetes as a whole and WG Naming have dealt with a number of issues and comments from trolls. We anticipate anyone using this document to guide their own work will receive the same kind of attention. 

Kubernetes is a large enough open source project that we mostly see [sea lions](http://wondermark.com/1k62/) (concern trolls), who seek to debate false concerns with us legitimately to use up our energy and time.

When possible, we work with our GitHub and other moderation teams to shut these down at the source and delete issues. 

In cases where it’s unclear whether the poster is a legitimate user or a troll, we direct the work back to them: because they’re clearly “legitimately interested” in this topic, we ask them to join us in the WG Naming mailing list, drafting a formal suggestion (attached to an email address and identity we can track) and suggesting replacement terminology. Most trolls do not want to put in the effort.


### Kudos

This work would not have come into shape without referencing the following resources freely available online. We thank the authors of these original documents for helping guide our thoughts on the topic:

- [APA Style Guide: General Principles for reducing bias](https://apastyle.apa.org/style-grammar-guidelines/bias-free-language/general-principles) 
- [Shopify Polaris Content Guidelines: Descriptive vs. Evocative names](https://polaris.shopify.com/content/naming#section-descriptive-vs-evocative-names)
- [CNET: Twitter engineers: out with the old words...](https://www.cnet.com/news/twitter-engineers-replace-racially-loaded-tech-terms-like-master-slave/) 
