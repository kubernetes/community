<!-- omit in toc -->
# Kubernetes Contributor Cheat Sheet

[English](README.md)

Questo documento fornisce una lista di risorse utili per contribuire a Kubernetes, suggerimenti, trucchi e
best practices utilizzate nel progetto Kubernetes. E' un "TL;DR" o un breve
condensato di informazioni utili per rendere migliore la tua esperienza di GitHub contributor.

**Sommario**

- [Helpful Resources](#helpful-resources)
  - [Getting Started](#getting-started)
  - [SIGs and Other Groups](#sigs-and-other-groups)
  - [Community](#community)
  - [Workflow](#workflow)
  - [Tests](#tests)
  - [Important Email Aliases](#important-email-aliases)
  - [Other Useful Links](#other-useful-links)
- [Communicating Effectively on GitHub](#communicating-effectively-on-github)
  - [How to be Excellent to Each Other](#how-to-be-excellent-to-each-other)
    - [Examples of Good/Bad Communication](#examples-of-goodbad-communication)
- [Submitting a Contribution](#submitting-a-contribution)
  - [Signing the CLA](#signing-the-cla)
  - [Opening and Responding to Issues](#opening-and-responding-to-issues)
    - [Creating an Issue](#creating-an-issue)
    - [Responding to an Issue](#responding-to-an-issue)
  - [Opening a Pull Request](#opening-a-pull-request)
    - [Creating a Pull Request](#creating-a-pull-request)
    - [Example PR Description](#example-pr-description)
    - [Troubleshooting a Pull Request](#troubleshooting-a-pull-request)
  - [Labels](#labels)
- [Working Locally](#working-locally)
  - [Branch Strategy](#branch-strategy)
    - [Adding Upstream](#adding-upstream)
    - [Keeping Your Fork in Sync](#keeping-your-fork-in-sync)
    - [Squashing Commits](#squashing-commits)

---

## Risorse utili

### Per cominciare

- [Contributor Guide] - Documento che descrive come cominciare a contribuire al Kubernetes project.
- [Developer Guide] - Documento che descrive come cominciare a contribuire codice al Kubernetes project.
- [Security and Disclosure Information] - Documento che descrive come segnalare vulnerabilities e come 
  viene gestito il security release process (hotfix).

### SIGs and Other Groups

- [Master Group List][sigs]

### Community

- [Calendar] - Visaluzza tutti gli eventi della Kubernetes Community (SIG/WG meetings,
  eventi etc.)
- [kubernetes-dev] - La mailing list su Kubernetes development
- [Kubernetes Forum] - Forum ufficale di Kubernetes.
- [Slack channels] - Slack channel ufficiale di Kubernetes.
- [Stack Overflow] - Il sito dove gli utenti di Kubernetes possono inviare le loro domande.
- [YouTube Channel] - YouTube channel ufficiale di Kubernetes.

### Workflow

- [Gubernator Dashboard] - Visualizza le Pull request che richiedono la tua attenzione.
- [Prow] - Kubernetes CI/CD System.
- [Tide] - Prow plugin che gestisce il merge delle PR e l'esecuzione dei tests. [Tide Dashboard]
- [Bot commands] - Guida ai comandi utilizzati per interagire con i bot di Kubernetes (examples:
  `/cc`, `/lgtm`, and `/retest`)
- [GitHub labels] - Elenco delle labels utilizzate nel Kubernetes Project
- [Kubernetes Code Search], mantenuto da [@dims]

### Tests

- [Prow] - Kubernetes CI/CD System.
- [Test Grid] - Visualizzo lo storico dei test e tutte le informazioni associate.
- [Triage Dashboard] - Aggrega problematiche simili per semplificare il troubleshooting.

### Important Email Aliases

- community@kubernetes.io - Invia una Mail ai membri del community team (SIG Contributor
  Experience) per discutere una issue sulla community.
- conduct@kubernetes.io - Contatta il Code of Conduct committee (private mailing list).
- github@kubernetes.io - Invia una Mail privata al [GitHub Administration Team], per argomenti sensibili.
- steering@kubernetes.io - Invia una Mail allo steering committee. Questa archivio di mail è pubblica.
- steering-private@kubernetes.io - Invia una provata Mail allo steering committee, per argomenti sensibili.
- social@cncf.io - Contatta il CNCF social team; blog, twitter account, e tutte le altre attività social
  correlate.

### Other Useful Links

- [Developer Statistics] - Visualizza le developer statistics per tutti i progetti gestiti da CNCF.
- [Kubernetes Patch Release] - Visualizza il calendario delle Kubernetes patch releases e i contatti
  del release team.

---

## Comunicare efficacement su GitHub

### Come essere rispettosi gli uni degli Altri

Come primo step, familiarizza con il [Code of Conduct].

#### Esempi di buone/cattive comunicazioni

Quando stai segnalando un problema, o chiedendo aiuto, esprimi cortesemente la tua richiesta:

  🙂 “X doesn’t compile when I do Y, do you have any suggestions?”

  😞 “X doesn’t work! Please fix it!”

Quando stai chiudendo una PR, fornisci una spiegazione cordiale del motivo per cui
questa non rispetta i criteri per essere accettata.

🙂 “I’m closing this PR because this feature can’t support the use case X. In
   it's proposed form, it would be a better to be implemented with Y tool. Thank
   you for working on this.”

😞 “Why isn’t this following the API conventions? This should be done elsewhere!”

---

## Contribuire

### Firmare la CLA

Prima di poter inviare qualsiasi contribuzione a Kubernetes, devi firmare la
[Contributor License Agreement(CLA)][cla]. Il progetto Kubernetes accetta
contribuzioni _solo_ se tu o la tua azienda hanno firmato la CLA.

Nel caso tu incontri problemi nel firmare la CLA, segui le [CLA
troubleshooting guidelines].

### Aprire of rispondere ad una Issues

Le issues in GitHub sono lo strumento principale per tracciare elementi come bug reports,
richieste di nuove feature, o report di altri problemi come i failing tests. Le issue
**non** sono intese come strumenti per gestire [user support requests]. Per queste ultime,
verifica la [troubleshooting guide], segnala il problema su [Stack Overflow] o
accedi al [Kubernetes forum].

**References:**

- [Labels]
- [Prow commands][commands]

#### Creare una Issue

- Utilizza l'issue template se disponibile. L'ulizzo corretto degli issue template
  aiuta gli altri contributors nel rispondere alla tua issue.
- Segui le istruzioni definite nell'issue template.
- Fornisci una descrizione accurata della issue che stai sollevando.
- Assegna le [labels] opportune. Se non sei sicure, il [k8s-ci-robot][prow] bot
  ([Kubernetes CI bot][prow]) risponderà alla tua issue aggiungendo le label necessarie
  affinche la tua issue sia opportunamente valutata.
- Utilizza con parsimonia i comandi [`/assign @<username>`][assign] o
  [`/cc @<username>`][cc]. La tua issue sarà valuta in modo più efficace applicando le
  le label opportune anzichè assegnando più persone alla tua issue.

#### Rispondere ad una Issue

- Quanndo inizi a lavorare su una issue, lascia un commento in modo che gli altri
  contributors sono informati della cosa e quindi si eviti di duplicare lavoro.
- Se per caso risolvi una tua vecchia issue per conto tuo, prima di chiuderla lascia
  un commento sulla issue in modo che possa aiutare altri utenti che incontreranno
  lo stesso problema.
- Includi riferimenti al altre PRs o issues (o altro materiale pubblicamente accessibile).
  Esempio: _"ref: #1234"_. Questo aiuta ad identificare altre attività correlate
  che sono in corso/già state completate altrove.

### Aprire una Pull Request

Le pull requests in GitHub (PR) sono il principale modo per contribuire codice, documentazione
o altre forme di lavoro nei git repository di Kubernetes.

**References:**

- [Labels]
- [Prow commands][commands]
- [Pull request process]
- [GitHub workflow]

#### Creare una Pull Request

- Segui le istruzioni del pull request template se disponibile. Questo aiuterà
  quelli che risponderanno alla tua PR.
- Se la PR è per un [trivial fix] come un broken link, un typo o un errore di grammatica,
  rivedi tutto il documento per verificare se ci sono altri errori dello stesso tipo.
  Non aprire diverse PRs per small fixes nello stesso documento.
- Fai riferimento a qualsiasi issue correlata alla tua PR a alle issue che la tua PR risolve.
- Evita di creare single commits con un numero elevato di modifiche. Invece, suddividi la tua PR
  in una serie di commit più semplici e logicamente consistenti. Questo semplifica
  il lavoro di review della tua PR.
- Aggiungi dei commenti alla tua PR dove ritieni che qualcosa necessiti di una spiegazione
  aggiuntiva.
- Utilizza con parsimonia il comando [`/assign @<username>`][assign].
  Avere PR con un numero elevato di reviewer non ti garantisce di avere una review più
  rapidamente.
- Se consideri la tua PR ancora un _"Work in progress"_ aggiungi il prefisso `[WIP]`
  al nome della PR oppure utilizza il comando [`/hold`][hold]. Questo impedisce all tua
  PR di fare merge fino a che `[WIP]` o hold sono rimossi.
- Se la tua PR non è stata rivista, non chiuderla e non aprirne un'altra con le
  stesse modifiche. Puoi pingare i tuoi reviewer con un commento dove con `@<github username>`.
- Se non ottieni risposta entro qualche giorno, posta un link della tua PR nel
  `#pr-reviews` channel su Slack per trovare altri reviewers.

#### Esempi di PR description

```
Ref. #3064 #3097
All files owned by SIG testing were moved from `/devel` to the new folder `/devel/sig-testing`.

/sig contributor-experience
/cc @stakeholder1 @stakeholder2
/kind cleanup
/area developer-guide
/assign @approver1 @approver2 @approver3
```

Cosa c'è in questa PR:

- **Line 1** - Il riferimento ad altre issues o PRs (#3064 #3097).
- **Line 2** - Una semplice descrizione di quali modifiche questa PR implementa.
- **Line 4** - L'indicazione del [SIG][sigs] di riferimento tramite il [comando][commands]
  `/sig contributor-experience`..
- **Line 5** - L'elenco dei reviewer che potrebbero essere interessati a questa PR è
  specificato utilizzando il comando [`/cc`][cc].
- **Line 6** - Il comando [`/kind cleanup`][kind] aggiunge una [label][labels] che
  categorizza la PR come attività di clean up di codice, processi, o riduzione del technical
  debt.
- **Line 7** - Il comando [`/area developer-guide`][kind] categorizza la PR
  come modifica nell'area della developer guide.
- **Line 8** - Il comando [`/assign`][assign] assegna uno o più reviewer alla PR.
  Un approver saà suggerito dal [k8s-ci-robot][prow] che seleziona dalle liste
  di owners definite negli [OWNERS] file. Gli approver aggiungeranno la label
  [`/approve`][approve] alla PR quando sarà pronta per il merge.

#### Troubleshooting di una Pull Request

Dopo che la tua PR è inviata, una serie di test è eseguita dalla piattaforma
di CI di Kubernetes, [Prow]. Se uno dei test fallisce, [k8s-ci-robot][prow]
aggiungerà un commento alla PR con il link ai test falliti e ai relativi logs.

Ogni nuova commit sulla tua PR comporta in automatico la ripetizione dei test.

In alcune situazioni ci possono anche essere dei problemi con la piattaforma
di CI di Kubernetes (flakes). Questi possono avvenire per svariate ragioni, anche se
gli stessi test passano localmente. Puoi ri-eseguire i test con il comando `/retest`.

Per maggiori informazioni sulla risoluzione dei problemi sui test, vedi [Testing Guide].

### Labels

Kubernetes utilizza le [labels] per fare triage e categorizzare sia le issues
che le pull request. Applicare le label opportune aiuta la tua issue o la tua
PR ad essere valutata in modo più effettivo.

**References:**

- [Labels]
- [Prow commands][commands]

Label utilizzate frequenetemente:

- [`/sig <sig name>`][kind] Assegna ad un [SIG][SIGs] l'ownership della issue o
  della PR.
- [`/area <area name>`][kind] Associa la issue o la PR ad una specifica
  [area][labels].
- [`/kind <category>`][kind] [Categorizes][labels] definisce il tipo di una issue o
  della PR..

---

## Lavora localmente

Prima di inviare una pull request è necessario fare iun pò di lavoro localmenente.
Se non sei esperto di git, l' [Atlassian git tutorial] è un buon punto di partenza.
Come alternativa, [Git magic] tutorial di Stanford è una soluzione disponbile in
diverse lingue.

**References:**

- [Atlassian git tutorial]
- [Git magic]
- [GitHub workflow]
- [Testing locally]
- [Developer guide]

### Branch Strategy

Il progetto Kubernetes utilizza il _"Fork and Pull"_ workflow che è tipico di
diversi progetti su GitHub. Utilizzando la terminologia git terms, il tuo personal fork
è definito come _"`origin`"_ mentre il progetto vero e proprio è chiamato _"`upstream`"_.
Per mantenere il tuo branch personale (`origin`) allineato con il progetto (`upstream`),
è necessario che quest'ultimo sia collegato nel tuo ambiente di lavoro locale.

#### Aggiungere Upstream

Aggiungere `upstream` come remote, e disabilitare la possibilità di fare push su quest'ultimo.

```
# replace <upstream git repo> with the upstream repo url
# example:
#  https://github.com/kubernetes/kubernetes.git
#  git@github.com/kubernetes/kubernetes.git

git remote add upstream <upstream git repo>
git remote set-url --push upstream no_push
```

Eseguendo il comando `git remote -v` è possibile verificare la corretta configurazione
dei remotes.

#### Mantenere il tuo Fork in Sync

Fai Fetch delle ultime modifiche da `upstream` e a segure _"rebase"_  delle stesse sul
tuo `master` branch locale. Questo permetterà di mantenere sincronizzati il repository locale
con l' `upstream` project. Quando hai finito, riscordati di fare Push delle modifiche locali
sul tuo `remote master`.

```
git fetch upstream
git checkout master
git rebase upstream/master
git push
```

Devi ripetere questa operazione ogni volta prima di creare un nuovo branch locale
per lavorare su una nuova feature o su una fix.

```
git checkout -b myfeature
```

#### Squashing Commits

L'obiettivo principale dell'operazione di [squashing commits] è quello di garantire che
la history delle commit documenti in modo chiaro l'elenco delle modifiche che sono state
fatte. Solitamente questa attività viene fatta come ultima fare del processi di review
un PR. Se non se sicuro se fare una squash dei tuoi commit, è solitamente preferibile
lasciare più commit e attendere che gli altri contributors che stanno facendo la review
della tua PR esprimano un giudizio in merito.

E' possbile utilizzare un interactive rebase per scegliere quali commits tenere e di quali
invece vuoi fare squash, e a seguire fai force push del tuo branch.

```
git rebase -i HEAD~3
...
git push --force
```

[contributor guide]: /contributors/guide/README.md
[developer guide]: /contributors/devel/README.md
[gubernator dashboard]: https://gubernator.k8s.io/pr
[prow]: https://prow.k8s.io
[tide]: http://git.k8s.io/test-infra/prow/cmd/tide/pr-authors.md
[tide dashboard]: https://prow.k8s.io/tide
[bot commands]: https://go.k8s.io/bot-commands
[gitHub labels]: https://go.k8s.io/github-labels
[Kubernetes Code Search]: https://cs.k8s.io/
[@dims]: https://github.com/dims
[calendar]: https://calendar.google.com/calendar/embed?src=calendar%40kubernetes.io
[kubernetes-dev]: https://groups.google.com/forum/#!forum/kubernetes-dev
[slack channels]: http://slack.k8s.io/
[Stack Overflow]: https://stackoverflow.com/questions/tagged/kubernetes
[youtube channel]: https://www.youtube.com/c/KubernetesCommunity/
[triage dashboard]: https://go.k8s.io/triage
[test grid]: https://testgrid.k8s.io
[developer statistics]: https://k8s.devstats.cncf.io
[code of conduct]: /code-of-conduct.md
[user support requests]: /contributors/guide/issue-triage.md#determine-if-its-a-support-request
[troubleshooting guide]: https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/
[kubernetes forum]: https://discuss.kubernetes.io/
[pull request process]: /contributors/guide/pull-requests.md
[github workflow]: /contributors/guide/github-workflow.md
[prow]: https://git.k8s.io/test-infra/prow#prow
[cla]: /CLA.md#how-do-i-sign
[cla troubleshooting guidelines]: /CLA.md#troubleshooting
[commands]: https://prow.k8s.io/command-help
[kind]: https://prow.k8s.io/command-help#kind
[cc]: https://prow.k8s.io/command-help#cc
[hold]: https://prow.k8s.io/command-help#hold
[assign]: https://prow.k8s.io/command-help#assign
[SIGs]: /sig-list.md
[testing guide]: /contributors/devel/sig-testing/testing.md
[labels]: https://git.k8s.io/test-infra/label_sync/labels.md
[trivial fix]: /contributors/guide/pull-requests.md#10-trivial-edits
[GitHub workflow]: /contributors/guide/github-workflow.md#3-branch
[squashing commits]: /contributors/guide/pull-requests.md#6-squashing-and-commit-titles
[owners]: /contributors/guide/owners.md
[testing locally]: /contributors/guide/README.md#testing
[Atlassian git tutorial]: https://www.atlassian.com/git/tutorials
[git magic]: http://www-cs-students.stanford.edu/~blynn/gitmagic/
[Security and Disclosure Information]: https://kubernetes.io/docs/reference/issues-security/security/
[approve]: https://prow.k8s.io/command-help#approve
[GitHub Administration Team]: /github-management#github-administration-team
[Kubernetes Patch Release]: https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md
