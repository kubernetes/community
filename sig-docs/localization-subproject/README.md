# Localization Subproject

The Kubernetes Localization Subproject is owned by [SIG Docs](https://github.com/kubernetes/community/tree/master/sig-docs) and run by the [Subproject leads](#leadership).

The Kubernetes documentation has been localized into several different languages. Each localization is owned by a team of individuals that speaks the language.  A localization team is a self-sufficient community responsible for prioritizing work, translating content, and reviewing and merging PRs into the kubernetes/website repo for their localization. 

The Localization Subproject's aim is to provide a place for localization teams to share ideas and standardize processes across all localizations. This page covers documentation, processes, and roles for the Kubernetes Localization teams and the Localization Subproject.

## Meetings

Regular Localization Meeting: The first Monday of the month at 15:00 UTC (biweekly). [Convert it to your timezone](http://www.thetimezoneconverter.com/?t=15:00&tz=UTC)

- [Meeting notes and Agenda](https://docs.google.com/document/d/1NwO1AN8Ea2zlK8uAdaDAKf1-LZDAFvSewIfrKqfl5No/edit#)
- [Meeting recordings](https://www.youtube.com/playlist?list=PL69nYSiGNLP3b5hlx0YV7Lo7DtckM84y8)

A localization team may also hold language specific meetings. If you are interested in attending a language specific meeting see the [Localization Teams](#localization-teams) section for details on how to contact each localization team.

## Subproject Leadership

- **Subproject lead:** [Abigail McCarthy](https://github.com/a-mccarthy)

- **Emeritus leads:** [Brad Topol](https://github.com/bradtopol)

## Contact

- Slack: [#sig-docs-localizations](https://kubernetes.slack.com/archives/C0191RDKHU1) Each localization also has a language specific channel, noted [below](#localization-teams). 
- Mailing list: https://groups.google.com/g/kubernetes-sig-docs-localization

## Responsibilities

The Localization Subproject lead is **responsible** for

- Leading Subproject meetings and taking notes.
- Driving collaborative discussions and decisions around localization processes and standards.
- Ensuring the communication of localization and SIG Docs policy changes to localization teams.
- Advocating for localization-friendly practices within SIG Docs and the broader Kubernetes community.
- Helping Localization teams understand SIG Docs processes, especially when languages are starting out
- Making sure that [Localization process documentation](https://kubernetes.io/docs/contribute/localization/) is kept up to date.

The Localization Subproject lead is **not responsible** for

- Running localization efforts for a language
- Dictating how localization teams do work or organize
- Making decisions for localizations policies or teams without the localization team's input

### Skills and Experience Required 

- Familiar with the [localization processes](https://kubernetes.io/docs/contribute/localization/).
- Have strong written and verbal communication skills. 
- A working knowledge of Github
- A working knowledge of kubernetes/website [review and merge processes](https://kubernetes.io/docs/contribute/review/).

Time commitment for a Localization lead is 1-2 hours per week.

# Localization teams

The following is a list of all localization teams. Localization teams are responsible for keeping this list up-to-date. 

| Language |  Leads | Docs Link | Slack Channel | Language specific meeting details (if available)|
|--|--|--|--|--|
| Arabic | | In progress | [#kubernetes-docs-ar](https://kubernetes.slack.com/archives/CP9FKRD51) |
| Bengali | Shahriyar Al Mustakim Mitul ([@mitul3737](https://github.com/mitul3737)) | In progress | [#kubernetes-docs-bn](https://kubernetes.slack.com/archives/CQ0TD298C) |
| Chinese | Qiming Teng ([@tengqm](https://github.com/tengqm)) | https://kubernetes.io/zh-cn/ | [#kubernetes-docs-zh](https://kubernetes.slack.com/archives/CE3LNFYJ1) | 
| French | Rémy Léone ([@remyleone](https://github.com/remyleone)) | https://kubernetes.io/fr/ | [#kubernetes-docs-fr](https://kubernetes.slack.com/archives/CG838BFT9) |
| German | | https://kubernetes.io/de/ | [#kubernetes-docs-de](https://kubernetes.slack.com/archives/CH4UJ2BAL) |
| Hindi | Anubhav Vardhan ([anubha-v-ardhan](https://github.com/anubha-v-ardhan)) | https://kubernetes.io/hi/ | [#kubernetes-docs-hi](https://kubernetes.slack.com/archives/CJ14B9BDJ) | 
| Indonesian | Aris Cahyadi Risdianto ([@ariscahyadi](https://github.com/ariscahyadi)) | https://kubernetes.io/id/ | [#kubernetes-docs-id](https://kubernetes.slack.com/archives/CJ1LUCUHM) |
| Italian | Fabrizio Pandini ([@fabriziopandini](https://github.com/fabriziopandini)) | https://kubernetes.io/it/ | [#kubernetes-docs-it](https://kubernetes.slack.com/archives/CGB1MCK7X) | 
| Japanese | Masahiro Kitamura ([@nasa9084](https://github.com/nasa9084)), Kohei Ota ([@inductor](https://github.com/inductor)) | https://kubernetes.io/ja/ | [#kubernetes-docs-ja](https://kubernetes.slack.com/archives/CAG2M83S8) | 
| Korean | Seokho Son ([@seokho-son](https://github.com/seokho-son)) | https://kubernetes.io/ko/ | [#kubernetes-docs-ko](https://kubernetes.slack.com/archives/CA1MMR86S) | 
| Portuguese | Ricardo Katz ([rikatz](https://github.com/rikatz)), Mauren Berti ([@stormqueen1990](https://github.com/stormqueen1990)), Edson C ([@edsoncelio](https://github.com/edsoncelio)) | https://kubernetes.io/pt-br/ | [#kubernetes-docs-pt](https://kubernetes.slack.com/archives/CJ21AS0NA) | 
| Russian | Dmitry Shurupov ([@shurup](https://github.com/shurup)) | https://kubernetes.io/ru/ | [#kubernetes-docs-ru](https://kubernetes.slack.com/archives/CPZ9KD9TN) |
| Spanish | Rael Garcia ([@raelga](https://github.com/raelga)), Victor Morales ([@electrocucaracha](https://github.com/electrocucaracha/)) | https://kubernetes.io/es/ | [#kubernetes-docs-es](https://kubernetes.slack.com/archives/CH7GB2E3B) | 
| Ukrainian | Maksym Vlasov ([@MaxymVlasov](https://github.com/MaxymVlasov)) | https://kubernetes.io/uk/ | [#kubernetes-docs-uk](https://kubernetes.slack.com/archives/CSKCYN138) | 
| Vietnamese | | https://kubernetes.io/vi/ | [#kubernetes-docs-vi](https://kubernetes.slack.com/archives/CPHAWNF1Q) |  

"In progress" teams have started localization work, but are not yet live on the Kubernetes site.

## Localization team structure

Anyone can start a new localization effort for any language or join an ongoing localization team. If you are interested in learning more about specific localization requirements or processes, see the [Localization Contributing Guide](https://kubernetes.io/docs/contribute/localization/).  

Localization teams are self-sufficient communities within SIG Docs and are responsible for

- Localizing content on kubernetes.io into their chosen language.  
- Timely reviews and merging PRs for their languages according to the [review and merge precesses](https://kubernetes.io/docs/contribute/review/) of the kubernetes/website repo.
- Building a community that can support their localization efforts. Note that you need at least two contributors to begin a localization team because contributors can't approve their pull requests. You will likely need more contributors, reviewers, and approvers to create a sustainable localization effort.
- Organizing and prioritizing work to localize content.
- Understanding and following the [Kubernetes Code of Conduct](https://kubernetes.io/community/code-of-conduct/) within their teams. Questions or concerns around Code of Conduct violations should be raised to the SIG Docs, Localization Subproject leadership, or [Kubernetes Code of Conduct Committee](https://github.com/kubernetes/community/tree/master/committee-code-of-conduct). 

The Localization Subproject leads and the SIG Docs leadership are always available to help localization teams by answering questions and providing feedback.

### Localization team leads

A leader of a localization team is someone who is responsible for helping to guide the localization efforts for a given language. Its recommended that each localization team has at least 2 leads.   

Some recommended skills for localization leads include, 
* Excellent written and verbal skills in the given language.
* Experience with the kubernetes/website repo localization [branching strategy](https://kubernetes.io/docs/contribute/localization/#branching-strategy).
* Experience reviewing English language PRs to develop an understanding the review processes and repo structure. Its recommended that leaders participate in the SIG Docs [PR Wrangler Shadow Program](https://kubernetes.io/docs/contribute/participate/pr-wranglers/#pr-wrangler-shadow-program).


## Localization processes

See the [Localization Contributing Guide](https://kubernetes.io/docs/contribute/localization/) for full information on localization processes and policies. This page includes information on [starting a new localization](https://kubernetes.io/docs/contribute/localization/#start-a-new-localization), [branching strategies](https://kubernetes.io/docs/contribute/localization/#branching-strategy) for localizing content, and more. 