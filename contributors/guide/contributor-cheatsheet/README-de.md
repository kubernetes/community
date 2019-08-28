# Cheat-Sheet für Kubernetes-Beitragende


Eine Liste von oft-genutzten Ressourcen zum Beitragen zu Kubernetes; Tipps, Tricks
und allgemeine Best-Practices innerhalb des Kubernetes-Projekts. 
Es ist als "TL;DR" (too long, didn't read - zu lang, nicht gelesen) oder Schnellreferenz
nützlicher Informationen gedacht, um deine Beitragserfahrung auf GitHub besser zu machen. 

**Inhaltsverzeichnis**
- [Hilfreiche Ressourcen](#Hilfreiche-Ressourcen)
  - [Wie man anfängt](#Wie-man-anfängt)
  - [SIGs und andere Gruppen](#SIGS-und-andere-Gruppen)
  - [Community](#Community)
  - [Important Email Aliases](#Important-Email-Aliases)
  - [Ablauf](#Ablauf)
  - [Tests](#Tests)
  - [Wichtige Email-Adressen](#Wichtige-Email-Adressen)
  - [Andere nützliche Links](#Andere-Nützliche-Links)
- [Effektive Kommunikation auf GitHub](#Effektive-Kommunikation-auf-GitHub)
  - [Seid exzellent zueinander](#Seid-exzellent-zueinander)
    - [Beispiele für gute/schlechte Kommunikation](#Beispiele-für-guteschlechte-Kommunikation)
- [Einen Beitrag einreichen](#Einen-Beitrag-einreichen)
  - [Das CLA unterzeichnen](#Das-CLA-unterzeichnen)
  - [Erstellen von und Antworten auf Issues](#Erstellen-von-und-Antworten-auf-Issues)
    - [Erstellen eines Issues](#Erstellen-eines-Issues)
    - [Antworten auf Issues](#Antworten-auf-Issues)
  - [Öffnen eines Pull-Requests](#Öffnen-eines-Pull-Requests)
    - [Erstellen eines Pull-Requests](#Erstellen-eines-Pull-Requests)
    - [Beispiel PR-Beschreibung](#Beispiel-PR-Beschreibung)
    - [Troubleshooting eines Pull-Requests](#Troubleshooting-eines-Pull-Requests)
  - [Labels](#Labels)
- [Lokal arbeiten](#Lokal-arbeiten)
  - [Branch-Strategie](#Branch-Strategie)
    - [Upstream-hinzufügen](#Upstream-hinzufügen)
    - [Synchron halten von Forks](#Synchron-halten-von-Forks)
  - [Commits squashen](#Commits-squashen)

---

## Hilfreiche Ressourcen

### Wie man anfängt

- [Contributor Guide] - Anleitung zum Beginn des Beitragens zum Kubernetes-Projekt.
- [Developer Guide] - Anleitung zum Beitragen von Code direkt zum Kubernetes-Projekt.
- [Security and Disclosure Information] - Anleitung zum Melden von Schwachstellen und 
  zur Sicherheit des Release-Prozesses.

### SIGs und andere Gruppen

- [Gruppenliste][SIGs]

### Community

- [Kalender] - Alle Kubernetes-Community-Events (SIG/WG Treffen, Events, usw.).
- [kubernetes-dev] - Kubernetesentwicklung-Mailingliste.
- [Kubernetes Forum] - Offizielles Kubernetes-Forum.
- [Slack Channels] - Offizieller Kubernetes-Slackchannel.
- [Stack-Overflow] - Zum Stellen von Kubernetes-Nutzerfragen.
- [YouTube Channel] - Offizieller Kubernetes-Youtubechannel.


### Ablauf

- [Gubernator Dashboard] - Eingehende und Ausgehende Pull-Requests, die Aufmerksamkeit erfordern.
- [Prow] - Kubernetes CI/CD-System.
- [Tide] - Prow-Plugin das Merges und Tests verwaltet. [Tide Dashboard]
- [Bot-Befehle] - Befehle zur Interaktion mit Kubernetes-Bots (zum Beispiel: `/cc`, `/lgtm`, und `/retest`)
- [GitHub Labels] - Liste an Labels des ganzen Kubernets-Projekts.
- [Kubernetes Code Search], von [@dims] maintained.


### Tests

- [Prow] - Kubernetes CI/CD-System.
- [Test Grid] - Ansicht historischer Tests und den zugehörigen Informationen.
- [Triage Dashboard] - Aggregiert ähnliche Fehler zur besseren Fehlerbehandlung.
- [Velodrome] - Dashboard zur Überwachung der Job- und Testgesundheit.


### Wichtige Email-Adressen

- community@kubernetes.io - Maile jemandem im Community-Team (SIG Contributor
  Experience) über ein Problem in der Community.
- conduct@kubernetes.io - Kontaktiere das Code-of-Conduct-Board, private Mailingliste.
- github@kubernetes.io - Maile dem [GitHub Administration Team] privat,
  für heikle Themen.
- steering@kubernetes.io - Maile der Leitungsgruppe. Öffentliche Adresse mit 
  öffentlichem Archiv.
- steering-private@kubernetes.io - Maile der Leitungsgruppe private, für heikle Themen.
- social@cncf.io - Kontaktiere das CNCF Socal Media Team; Blog, Twitteraccount, und
  andere soziale Bereiche.


### Andere nützliche Links

- [Developer Statistiken] - Entwicklerstatistiken für alle von CNCF gemanagten Projekte.

---

## Effektive Kommunikation auf GitHub


### Seid exzellent zueinander

Als ersten Schritt sollte man sich mit dem [Code of Conduct] bekannt machen. 


#### Beispiele für gute/schlechte Kommunikation

Wenn man ein Problem anspricht, oder Hilfe sucht, sollte man die Anfrage höflich
formulieren:

  🙂 “X kompiliert nicht, wenn ich Y mache, habt ihr Vorschläge dazu?

  😞 “X geht nicht, fixt das!" 

Beim Schließen eines PRs ist es sinnvoll eine freundliche und erklärende Nachricht
hinzuzufügen weshalb dieser nicht die Mergeanforderungen erfüllt.

🙂 “Ich schließe diesen PR, da dieses Feature nicht den Use-Case X unterstützt. In
   der forgeschlagenen Form wäre es besser dies mit Y umzusetzen. Danke für deine 
   Mitarbeit daran." 

😞 “Warum folgt das nicht den API-Konventionen? Das sollte woanders gemacht werden!"

---

## Einen Beitrag einreichen

### Das CLA unterzeichnen

Bevor man einen Beitrag einreichen kann, muss das [Contributor License Agreement (CLA) unterzeichnet][cla] 
werden. Das Kubernetes-Projekt kann _nur_ einen Beitrag annehmen, wenn du oder deine Firma 
das CLA unterzeichnet haben.

Solltest du Probleme beim Unterzeichnen des CLAs haben, folge den [CLA Troubleshooting Guidelines][CLA
troubleshooting guidelines].


### Erstellen von und Antworten auf Issues

GitHub Issues sind der meistgenutzte Weg um Dinge wie Bugreports und Enhancementrequests
zu verfolgen, oder andere Probleme wie fehlschlagende Tests zu melden.
Sie sind **nicht** als [Anfragen für Usersupport][user support requests] gedacht. Für diese
gibt es den [Troubleshooting Guide][troubleshooting guide], [Stack-Overflow] oder das
[Kubernetes-Forum][Kubernetes Forum].

**Referenzen:**
- [Labels]
- [Prow Commands][commands]


#### Erstellen eines Issues

- Nutze ein Issue-Template, wenn eines zur Verfügung steht. Das korrekte Template
  hilft anderen Beitragenden auf deinen Issue zu antworten.
  - Folge allen Anweisungen im Template selbst.
- Beschreibe den Issue detailliert.
- Füge sinnvolle [Labels][Labels] hinzu. Wenn du dir nicht sicher bist, hilft
  der [K8s-ci-robot][Prow] Bot ([Kubernetes CI bot][Prow]) mit Antworten auf Issues,
  damit diese effektive verwaltet werden. 
- Sei selektiv beim Zuweisen von Issues zu Nutzern mit [`/assign @<username>`][assign] 
  oder [`/cc @<username>`][cc]. Der Issue wird effektiver verwaltet, wenn sinnvolle
  Labels zugewiesen werden und weniger Menschen.


#### Antworten auf Issues

- Wenn du an einem Issue arbeitest, hinterlasse einen Kommentar damit andere
  wissen, dass du daran arbeitest und doppelte Arbeit vermieden wird.
- Falls du selbst eine Lösung findest, kommentiere den Issue mit einer Erklärung
  bevor du ihn schließt.
- Referenzen auf andere PRs und Issues (oder andere Materialen) mit zum Beispiel: 
  _"ref: #1234"_ sind sinnvoll. Diese helfen Bezüge darauf zu finden, die an
  anderer Stelle bearbeitet wurden.


### Öffnen eines Pull-Requests

Pull-Requests (PR) sind die meistgenutzte Variante um Code, Dokumentation oder andere Arten
von Arbeit beizutragen, die in einem Git-Repository gespeichert werden können.

**Referenzen:**
- [Labels]
- [Prow Commands][commands]
- [Pull-Request Prozess]
- [GitHub Workflow]


#### Erstellen eines Pull-Requests

- Folge den Anweisungen des Pull-Request-Templates, falls eines zur Verfügung steht.
  Es wird denen helfen, die auf den PR antworten.
- Für [triviale Fixes][trivial fix] wie kaputte Links, Schreib- oder Grammatikfehler
  durchsuche das gesamte Dokument für andere potentielle Fehler. Mehrere PRs für kleine
  Fixes im gleichen Dokument sind unnötig.
- Füge Referenzen auf verwandte Issues oder Issues die der PR lösen könnte hinzu.
- Vermeide unnötig große Änderungen in einem einzelnen Commit. Zerteile stattdessen
  den PR in mehrere kleine, logische Commits. Das erleichter Reviews.
- Kommentiere deinen eigenen PR wenn du meinst weitere Erklärungen sind nötig.
- Sei selektiv beim Erstellen des PRs mit [`/assign @<username>`][assign].
  Zu viele Reviewer garantieren keinen schnelleren PR-Review.
- Wenn der PR als _"Work in progress"_ angesehen wird, füge ein `[WIP]` am Anfang des
  Titels hinzu oder nutze den [`/hold`][hold] Befehl. Das stellt sicher, dass der PR 
  nicht gemerged wird bis das `[WIP]` oder hold aufgehoben worden sind.
- Falls dein PR noch nicht reviewed wurde, bitte schließe ihn nicht und öffne einen Neuen 
  mit den gleichen Änderungen. Pinge deine Reviewer stattdessen in einem Kommentar mit 
  `@<github username>` an.


#### Beispiel PR-Beschreibung

```
Ref. #3064 #3097
Alle Dateien im Besitz von SIG testing wurden von `/devel` zum neuen Ordner `/devel/sig-testing`
verschoben.

/sig contributor-experience
/cc @stakeholder1 @stakeholder2
/kind cleanup
/area developer-guide
/assign @approver1 @approver2 @approver3
```

Was steht in diesem PR:
- **Zeile 1** - Referenzen auf andere Issues oder PRs (#3064 #3097).
- **Zeile 2** - Eine kurze Beschreibung was in diesem PR getan wird.
- **Zeile 4** - [SIG][SIGs] Zuweisung mit dem [Befehl][commands]
  `/sig contributor-experience`..
- **Zeile 5** - Reviewer die Interesse an diesem spezifischen Issue oder PR 
  haben könnten, werden mit [`/cc`][cc] markiert.
- **Zeile 6** - Der [`/kind cleanup`][kind] Befehl fügt ein [Label][Labels] hinzu, das
  Issues oder PRs als Aufräumen von Code, Prozessen oder technologischer Schuld kategorisiert.
- **Zeile 7** - Der [`/area developer-guide`][kind] Befehl kategorisiert Issues oder PRs im Bezug
  zum Developerguide. 
- **Zeile 9** - Der Befehl [`/assign`][assign] fügt einen Approver zum PR hinzu. Ein Approver
  wird vom [k8s-ci-robot][Prow] vorgeschlagen und wird von der Liste der Eingentümer in der
  [OWNERS] Datei ausgewählt. file. Der Approver fügt das [`/approve`][approve] Label zum PR
  hinzu nachdem dieser reviewed wurde.


#### Troubleshooting eines Pull-Requests

Nachdem ein PR vorgeschlagen wurde, wird eine Reihe an Test von der Kubernetes
CL-Plattform [Prow] ausgeführt. Fall einer der Tests fehlschlägt, antwortet
der [K8s-ci-robot][Prow] auf den PR mit Links zu den fehlgeschlagenen Tests und
verfügbaren Logs.

Neue Commits zu diesem PR sorgen dafür dass diese Tests erneut ausgeführt werden.

Gelegentlich gibt es Probleme mit der Kubernetes CI-Plattform. Diese können aus 
verschiedenen Gründen auftreten, selbst wenn der Beitrag alle lokalen Tests besteht.
Man kann einen neuen Testdurchlauf mit dem `/retest` Befehl auslösen. 

Für mehr Informationen zum Troubleshooting von spezifischen Test, siehe den [Testing Guide].


### Labels

Kubernetes nutzt [Labels][Labels] zum sichten und kategorisieren von Issues und Pull Requests.
Die richtigen Labels helfen dabei den Issue oder PR sinnvoller einzusortieren.

**Referenzen:**
- [Labels]
- [Prow Commands][commands]

Oft genutzte Labels:
- [`/sig <sig name>`][kind] Fügt eine [SIG][SIGs] als Eigentümer des Issues oder PRs hinzu.
- [`/area <area name>`][kind] Verknüpft den Issue oder PR mit einem bestimmten [Bereich][labels].
- [`/kind <category>`][kind] [Kategorisiert][labels] den Issue oder PR.

---

## Lokal arbeiten

Bevor man einen Pull-Request vorschlagen kann, muss man lokal etwas Arbeit leisten.
Für Neueinsteiger zu git ist das [Atlassian Git-Tutorial][Atlassian git tutorial] 
ein guter Einstiegspunkt. Alternativ ist das [Git Magic Tutorical][Git magic] der
Standford Uni eine gute multi-linguale Option.

**Referenzen:**
- [Atlassion Git-Tutorial][Atlassian git tutorial]
- [Git magic]
- [GitHub Workflow]
- [Lokakes Testen][Testing locally]
- [Developer guide]


### Branch-Strategie

Das Kubernetes-Projekt nutzt den Standardworkflow von GitHub, der sich _"Fork and Pull"_ 
(auf Deutsch "abzweigen und ziehen") nennt.
In Begriffen aus der Git-Welt wird dein persönlicher Fork als _"`origin`"_ 
(auf Deutsch "Ursprung") und das eigentliche Projekt-Gitrepository als _"`upstream`"_
(wörtlich "flussaufwärts") bezeichnet. 
Um deinen persönlichen Branch (`origin`) mit dem Projekt (`upstream`) aktuell zu halten,
muss das innerhalb der lokalen Kopie konfiguriert werden.


#### Upstream hinzufügen

Füge `upstream` als sogenanntes remote hinzu und konfiguriere es so, dass man nicht dorthin
pushen kann.

```
# Ersetze <upstream git repo> mit der Upstreamrepo-URL
# Beispiel:
#  https://github.com/kubernetes/kubernetes.git
#  git@github.com/kubernetes/kubernetes.git

git remote add upstream <upstream git repo>
git remote set-url --push upstream no_push
```

Das kann via `git remote -v` verifiziert werden, indem alle konfigurierten Remote-Repos 
aufgelistet werden.


#### Synchron halten von Forks

Hole alle Änderungen von `upstream` ab und _"rebase"_ diese auf deinem lokalen
`Master` Branch. Das wird dein lokales Repo mit dem `upstream` Projekt synchronisieren.

```
git fetch upstream
git checkout master
git rebase upstream/master
```

Das sollte minimal bevor der Erstellung eines neuen Branches für ein Feature oder
einen Fix passieren. 

```
git checkout -b myfeature
```

#### Commits squashen

Der Hauptzweck von [Commits squashen]("Commits zerquetschen") ist die Erstellung
einer sauberen, lesbaren Githistorie oder eines Logs der Änderungen die gemacht wurden.
Normal wird das in der letzten Phase einer PR Revision getan. Wenn du dir unsicher bist,
ob du deine Commits squashen solltest, ist es besser mehrere zu lassen und es dem Urteil
anderer Beitragenden zu überlassen, die als Reviewer und Approver für den PR zugeteilt wurden.


[Contributor Guide]: /contributors/guide/README.md
[Developer Guide]: /contributors/devel/README.md
[Gubernator Dashboard]: https://gubernator.k8s.io/pr
[Prow]: https://prow.k8s.io
[Tide]: http://git.k8s.io/test-infra/prow/cmd/tide/pr-authors.md
[Tide Dashboard]: https://prow.k8s.io/tide
[Bot-Befehle]: https://go.k8s.io/bot-commands
[GitHub Labels]: https://go.k8s.io/github-labels
[Kubernetes Code Search]: https://cs.k8s.io/
[@dims]: https://github.com/dims
[Kalender]: https://calendar.google.com/calendar/embed?src=cgnt364vd8s86hr2phapfjc6uk%40group.calendar.google.com
[kubernetes-dev]: https://groups.google.com/forum/#!forum/kubernetes-dev
[Slack Channels]: http://slack.k8s.io/
[Stack-Overflow]: https://stackoverflow.com/questions/tagged/kubernetes
[Youtube Channel]: https://www.youtube.com/c/KubernetesCommunity/
[Triage Dashboard]: https://go.k8s.io/triage
[Test Grid]: https://testgrid.k8s.io
[Velodrome]: https://go.k8s.io/test-health
[Developer Statistiken]: https://k8s.devstats.cncf.io
[Code of Conduct]: /code-of-conduct.md
[user support requests]: /contributors/guide/issue-triage.md#determine-if-its-a-support-request
[troubleshooting guide]: https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/
[Kubernetes Forum]: https://discuss.kubernetes.io/
[Pull-Request Prozess]: /contributors/guide/pull-requests.md
[GitHub Workflow]: /contributors/guide/github-workflow.md
[Prow]: https://git.k8s.io/test-infra/prow#prow
[cla]: /CLA.md#how-do-i-sign
[CLA troubleshooting guidelines]: /CLA.md#troubleshooting
[commands]: https://prow.k8s.io/command-help
[kind]: https://prow.k8s.io/command-help#kind
[cc]: https://prow.k8s.io/command-help#cc
[hold]: https://prow.k8s.io/command-help#hold
[assign]: https://prow.k8s.io/command-help#assign
[SIGs]: /sig-list.md
[Testing Guide]: /contributors/devel/sig-testing/testing.md
[Labels]: https://git.k8s.io/test-infra/label_sync/labels.md
[trivial fix]: /contributors/guide/pull-requests.md#10-trivial-edits
[GitHub Workflow]: /contributors/guide/github-workflow.md#3-branch
[Commits squashen]: /contributors/guide/pull-requests.md#6-squashing-and-commit-titles
[OWNERS]: /contributors/guide/owners.md
[Testing locally]: /contributors/guide/README.md#testing
[Atlassian git tutorial]: https://www.atlassian.com/git/tutorials
[Git magic]: http://www-cs-students.stanford.edu/~blynn/gitmagic/
[Security and Disclosure Information]: https://kubernetes.io/docs/reference/issues-security/security/
[approve]: https://prow.k8s.io/command-help#approve
[GitHub Administration Team]: /github-management#github-administration-team
