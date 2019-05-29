# Kubernetes Contributor Cheat Sheet

Un elenco di risorse comuni quando contribuisci a Kubernetes, consigli, trucchi e
best practice comuni utilizzate all'interno del progetto Kubernetes. È un "TL; DR" o
riferimento rapido di informazioni utili per rendere la tua esperienza di contributo GitHub
meglio.

**Sommario**
- [Risorse utili](#risorse utili)
    - [Guida introduttiva](#Guida introduttiva)
    - [SIG e altri gruppi](#SIGS-e-Altri-Gruppi)
    - [Community](#Community)
    - [Alias ​​email importanti](#Importanti-Alias-email)
    - [Flusso di lavoro](#flusso di lavoro)
    - [Test](#test)
    - [Altri collegamenti utili](#Altri collegamenti utili)
- [Comunicare efficacemente su GitHub](#Comunicare-efficacemente-su-GitHub)
    - [Come essere eccellenti l'un l'altro](#Come essere eccellenti l'un l'altro)
        - [Esempi di comunicazione buona / cattiva](#esempi-di-comunicazione GoodBad)
- [Invio di un contributo](#Invio di un contributo)
    - [Firma del CLA](#signing-the-CLA)
    - [Apertura e risposta ai problemi](#Apertura e risposta ai problemi)
        - [Creazione di un problema](#Creazione di un problema)
        - [Risposta a un problema](#risposta a un problema)
    - [Apertura di una richiesta di pull](#Apertura di una richiesta di pull)
        - [Creazione di una richiesta di pull](#Creazione di una richiesta di pull)
        - [Esempio Descrizione PR](#Esempio-Descrizione-PR)
        - [Risoluzione dei problemi di una richiesta di pull](#Risoluzione dei problemi-a-Pull-Richiesta)
    - [Etichette](#etichette)
- [Lavorare localmente](#Lavorare-Localmente)
    - [Strategia delle filiali](#Strategia delle filiali)
        - [Aggiunta Upstream](#Aggiunta-Upstream)
        - [Mantenendo il fork sincronizzato](#Mantenendo il fork sincronizzato)
    - [Squashing Commits](#Squashing-Commits)
---

## Risorse utili

### Iniziare

- [Contributor Guide] - Guida su come iniziare a contribuire a Kubernetes
  Progetto.
- [Guida per lo sviluppatore] - Guida per contribuire con il codice direttamente ai Kubernetes
  Progetto.
- [Informazioni sulla sicurezza e divulgazione] - Guida per segnalare le vulnerabilità
  e il processo di rilascio della sicurezza.

### SIG e altri gruppi

- [Elenco gruppi principali] [sigs]

### Comunità

- [Calendario] - Visualizza tutti gli eventi della community di Kubernetes (riunioni SIG / WG,
  eventi ecc.)
- [kubernetes-dev] - La mailing list di sviluppo di Kubernetes
- [Forum di Kubernetes] - Forum ufficiale di Kubernetes.
- [Slack channels] - Official Kubernetes Slack.
- [StackOverflow] - Un posto dove porre le domande per gli utenti finali di Kubernetes.
- [Canale YouTube] - Canale ufficiale per la comunità di Kubernetes.


### Flusso di lavoro

- [Gubernator Dashboard] - Visualizza le richieste di pull in entrata e in uscita che richiedono
  la tua attenzione.
- [Prow] - Sistema Kubernetes CI / CD.
- [Tide] - Plugin Prow che gestisce unioni e test. [Tide Dashboard]
- [Comandi Bot] - Comandi utilizzati per interagire con Kubernetes Bots (esempi:
  `/ cc`,` / lgtm` e `/ retest`)
- [Etichette GitHub] - Elenco delle etichette utilizzate in tutto il Progetto Kubernetes
- [Ricerca codice Kubernetes], gestita da [@dims]


### Test

- [Prow] - Sistema Kubernetes CI / CD.
- [Test Grid] - Visualizza i test storici e le informazioni associate.
- [Triage Dashboard] - Aggrega guasti simili insieme per migliorare
  risoluzione dei problemi.
- [Velodrome] - Dashboard per tracciare il lavoro e testare la salute.


### Alias ​​email importanti

- community@kubernetes.io - Posta qualcuno sul team della comunità (collaboratore SIG
  Esperienza) su un problema di comunità.
- conduct@kubernetes.io - Contatta il comitato del codice di condotta, mailing privato
  elenco.
- steering@kubernetes.io - Scrivi al comitato direttivo. Indirizzo pubblico con
  archivio pubblico.
- steering-private@kubernetes.io - Invia per e-mail il comitato direttivo, per
  articoli sensibili
- social@cncf.io - Contatta il team sociale di CNCF; blog, account twitter e
  altre proprietà sociali.


### Altri link utili

- [Statistiche sviluppatore] - Visualizza le statistiche degli sviluppatori per tutto il CNCF gestito
  progetti.

---

## Comunicare in modo efficace su GitHub


### Come essere eccellenti gli uni con gli altri

Come primo passo, familiarizza con il [Codice di condotta].


#### Esempi di comunicazione buona / cattiva

Quando sollevi un problema o cerchi assistenza, sii gentile con la tua richiesta:

   🙂 "X non si compila quando faccio Y, hai qualche suggerimento?"

   😞 "X non funziona! Per favore aggiustalo! "

Quando si chiude un PR, trasmettere un messaggio esplicativo e cordiale che spiega
perché non soddisfa i requisiti per essere uniti.

I "Chiudo questo PR perché questa funzione non supporta il caso d'uso X. In
    è una forma proposta, sarebbe meglio implementarla con lo strumento Y. grazie
     per aver lavorato su questo. "

Why "Perché non segue le convenzioni API? Questo dovrebbe essere fatto altrove! "

---

### Firma del CLA

Prima di poter inviare un contributo, devi [firmare la Licenza di Contributor
Accordo (CLA)] [cla]. Il progetto Kubernetes può accettare solo un contributo
se tu o la tua azienda avete firmato il CLA.

In caso di problemi durante la firma del CLA, seguire la [CLA
linee guida sulla risoluzione dei problemi].


### Aprire e rispondere ai problemi

I problemi di GitHub sono il mezzo principale per tenere traccia di cose come segnalazioni di bug,
richieste di miglioramento o segnalazione di altri problemi, come test non riusciti. Loro sono
** non ** destinato a [richieste di supporto utente]. Per quelli, per favore controlla con il
[guida alla risoluzione dei problemi], segnala il problema a [Stack Overflow] o follow up su
il [forum di Kubernetes].

**Riferimenti:**
- [Etichette]
- [Prow commands] [comandi]


#### Creare un problema

- Utilizza un modello di problema se disponibile. Usando quello corretto aiuterà l'altro
  contributori nel rispondere al tuo problema.
  - Segui le istruzioni descritte nel modello di problema stesso.
- Sii descrittivo del problema che stai sollevando.
- Assegna [etichette] appropriate. Se non sei sicuro, il [k8s-ci-robot] [prow] bot
  ([Kubernetes CI bot] [prow]) risponderà al tuo problema con le etichette necessarie
  per essere efficacemente triaged.
- Sii selettivo quando assegni i problemi usando [`/ assign @ <username>`] [assegna] o
  [`/ cc @ <nome utente>`] [cc]. Il tuo problema verrà valutato in modo più efficace
  correggere le etichette assegnando più persone al problema.


#### Risposta a un problema

- Quando affronti un problema, commentalo facendo sapere agli altri che stai lavorando
  per evitare il doppio lavoro.
- Quando hai risolto qualcosa da solo in qualsiasi momento, commenta
  il problema di far conoscere la gente prima di chiuderla.
- Includere riferimenti ad altri PR o problemi (o qualsiasi materiale accessibile),
  esempio: _ "ref: # 1234" _. È utile identificare che lavoro correlato è stato
  indirizzato da qualche altra parte.


### Aprire una richiesta di pull

Le richieste di pull (PR) sono il mezzo principale per fornire il codice, la documentazione o
altre forme di lavoro che verrebbero archiviate all'interno di un repository git.

**Riferimenti:**
- [Etichette]
- [Prow commands] [comandi]
- [Processo di richiesta di pull]
- [flusso di lavoro Github]


#### Creare una richiesta di pull

- Seguire le indicazioni del modello di richiesta di pull, se disponibile. esso
  aiuterà quelli che rispondono al tuo PR.
- Se una [correzione banale] come un collegamento interrotto, un errore di digitazione o di grammatica, rivedere il
  intero documento per altri potenziali errori. Non aprire più PR per
  piccole correzioni nello stesso documento.
- Riferisci eventuali problemi relativi al tuo PR o problemi che potrebbero essere risolti da PR.
- Evita di creare modifiche troppo grandi in un singolo commit. Invece, rompere il tuo PR
  in più piccoli commit logici. Questo rende più facile per il tuo PR essere
  rivisto.
- Commenta il tuo PR in cui credi che qualcosa potrebbe essere necessario ulteriormente
  spiegazione.
- Sii selettivo quando assegni il tuo PR con [`/ assign @ <username>`] [assegna].
  Assegnare revisori eccessivi non produrrà una revisione PR più rapida.
- Se il tuo PR è considerato un _ "Work in progress" _ prefisso il nome con `[WIP]`
  oppure usare il comando [`/ hold`] [hold]. Ciò impedirà la fusione del PR
  fino a quando `[WIP]` o hold viene sollevato.
- Se non hai avuto la tua recensione pubblicitaria, non chiudere e aprire un nuovo PR con il
  stessi cambiamenti. Esegui il ping dei revisori in un commento con `@ <nomeutente github>`.


#### Esempio PR Descrizione

`` `
Ref. 3064 # 3097
Tutti i file di proprietà del test SIG sono stati spostati da `/ devel` alla nuova cartella` / devel / sig-testing`.

/ sig contributor-experience
/ cc @ stakeholder1 @ stakeholder2
/ tipo di pulizia
/ area developer-guide
/ Assegna @ approver1 @ approvatore2 @ approvatore3
`` `

Cosa c'è in quella PR:
- ** Linea 1 ** - Riferimento ad altri numeri o PR (# 3064 # 3097).
- ** Linea 2 ** - Una breve descrizione di ciò che viene fatto nel PR.
- ** Linea 4 ** - [SIG] [sigs] assegnazione con il [comando] [comandi]
  `/ sig contributor-experience` ..
- ** Linea 5 ** - I revisori che potrebbero avere interesse su questo specifico problema o PR sono
  specificato con il comando [`/ cc`] [cc].
- ** Linea 6 ** - Il comando [`/ kind cleanup`] [kind] aggiunge una [etichetta] [etichette]
  classifica il problema o il PR come correlato alla pulizia del codice, del processo o tecnico
  debito.
- ** Linea 7 ** - Il comando [`/ area developer-guide`] [kind] classifica il problema o
  PR correlato alla guida dello sviluppatore.
- ** Linea 8 ** - Il comando [`/ assign`] [assegna] assegna un approvatore al PR.
  Un approver sarà suggerito dal [k8s-ci-robot] [prua] ed è selezionato da
  l'elenco dei proprietari nel file [PROPRIETARI]. Aggiungeranno il
  [`/ approve`] [approva] etichetta al PR dopo che è stata rivista.



#### Risoluzione dei problemi di una richiesta di pull

Dopo la proposta del PR, una serie di test viene eseguita dall'IC di Kubernetes
piattaforma, [Prow]. Se uno dei test fallisce, il [k8s-ci-robot] [prow]
risponderà al PR con link ai test falliti e ai log disponibili.

Spingendo nuovi commit al tuo PR, i test verranno automaticamente riavviati.

Occasionalmente possono esserci problemi con la piattaforma CI di Kubernetes. Questi possono accadere
per una vasta gamma di motivi, anche se il tuo contributo passa a livello locale
test. È possibile attivare una nuova esecuzione dei test con il comando `/ retest`.

Per ulteriori informazioni sulla risoluzione di test specifici, consultare la [Guida al test].


### Etichette

Kubernetes usa [etichette] per categorizzare e triage problemi e Pull Requests.
Applicare le etichette giuste aiuterà di più il tuo problema o le tue PR
in modo efficace.

**Riferimenti:**
- [Etichette]
- [Prow commands] [comandi]

Etichette usate frequentemente:
- [`/ sig <sig name>`] [tipo] Assegna un [SIG] [SIGs] alla proprietà del problema
  o PR.
- [`/ area <nome area>`] [tipo] Associa il problema o le PR a uno specifico
  [AREA] [] etichette.
- [`/ kind <category>`] [kind] [Categorizes] [labels] il problema o PR.

---

## Lavorare localmente

Prima di proporre una richiesta di pull, dovrai svolgere un certo livello di lavoro
localmente. Se sei nuovo a git, il [tutorial git Atlassian] è un buon inizio
punto. In alternativa, il tutorial di [Git magic] di Stanford è un bene
opzione multi-lingua.

**Riferimenti:**
- [Atlitian git tutorial]
- [Git magic]
- [flusso di lavoro Github]
- [Testing localmente]
- [Guida per gli sviluppatori]


### Strategia del ramo

Il progetto Kubernetes utilizza un flusso di lavoro _ "Fork and Pull" _ standard
GitHub. In termini git, la tua forcella personale viene chiamata _ "` origine` "_ e
il vero repository git del progetto è chiamato _ "` upstream` "_. Per mantenere il tuo
ramo personale (`origine`) aggiornato con il progetto (` upstream`), deve essere
configurato all'interno della tua copia di lavoro locale.


#### Aggiunta di Upstream

Aggiungi `upstream` come remoto e configuralo in modo da non poterlo spingere.

`` `
# sostituisce <upstream git repo> con l'url repo upstream
# esempio:
# https://github.com/kubernetes/kubernetes.git
# git@github.com/kubernetes/kubernetes.git

git remote add upstream <upstream git repo>
git remote set-url --push upstream no_push
`` `

Questo può essere verificato eseguendo `git remote -v` che elencherà i tuoi configurati
telecomandi.


#### Mantenere la tua forcella in sincronizzazione

Recupera tutte le modifiche da `upstream` e _" rebase "_ sul tuo` master` locale
ramo. Questo sincronizzerà il repository locale con il progetto `upstream`.

`` `
git fetch upstream
git checkout master
git rebase upstream / master
`` `

Dovresti fare questo minimamente prima di creare un nuovo ramo su cui lavorare
funzionalità o correzione.

`` `
git checkout -b myfeature
`` `

#### Impegni di schiacciamento

Lo scopo principale di [squashing commits] è creare un git pulito e leggibile
cronologia o registro delle modifiche apportate. Di solito questo è fatto per ultimo
fase di una revisione di pubbliche relazioni. Se non sei sicuro se dovresti schiacciare i tuoi commit, lo fai
è meglio sbagliare dalla parte di avere di più e lasciarlo al giudizio di
gli altri contributori assegnati per esaminare e approvare il tuo PR.



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
[calendar]: https://calendar.google.com/calendar/embed?src=cgnt364vd8s86hr2phapfjc6uk%40group.calendar.google.com
[kubernetes-dev]: https://groups.google.com/forum/#!forum/kubernetes-dev
[slack channels]: http://slack.k8s.io/
[stackOverflow]: https://stackoverflow.com/questions/tagged/kubernetes
[youtube channel]: https://www.youtube.com/c/KubernetesCommunity/
[triage dashboard]: https://go.k8s.io/triage
[test grid]: https://testgrid.k8s.io
[velodrome]: https://go.k8s.io/test-health
[developer statistics]: https://k8s.devstats.cncf.io
[code of conduct]: /code-of-conduct.md
[user support request]: /contributors/guide/issue-triage.md#determine-if-its-a-support-request
[troubleshooting guide]: https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/
[stack overflow]: https://stackoverflow.com/questions/tagged/kubernetes
[kubernetes forum]: https://discuss.kubernetes.io/
[pull request process]: /contributors/guide/pull-requests.md
[github workflow]: /contributors/guide/github-workflow.md
[prow]: https://git.k8s.io/test-infra/prow#prow
[cla]: /CLA.md#how-do-i-sign
[cla troubleshooting guidelines]: /CLA.md#troubleshooting
[commands]: https://prow.k8s.io/command-help
[kind]: https://prow.k8s.io/command-help#kind
[cc]: https://prow.k8s.io/command-help#hold
[hold]: https://prow.k8s.io/command-help#hold
[assign]: https://prow.k8s.io/command-help#assign
[SIGs]: /sig-list.md
[testing guide]: /contributors/devel/sig-testing/testing.md
[labels]: https://git.k8s.io/test-infra/label_sync/labels.md
[trivial fix]: /contributors/guide/pull-requests.md#10-trivial-edits
[Github workflow]: /contributors/guide/github-workflow.md#3-branch
[squashing commits]: /contributors/guide/pull-requests.md#6-squashing-and-commit-titles
[owners]: /contributors/guide/owners.md
[testing locally]: /contributors/guide/README.md#testing
[developer guide]: /contributors/devel/README.md
[Atlassian git tutorial]: https://www.atlassian.com/git/tutorials
[git magic]: http://www-cs-students.stanford.edu/~blynn/gitmagic/
[Security and Disclosure Information]: https://kubernetes.io/docs/reference/issues-security/security/
