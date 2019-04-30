# Cheat Sheet pour contributeur Kubernetes

Une liste des ressources communes pour contribuer à Kubernetes, des trucs, des astuces et des bonnes pratiques communes utilisées dans le projet Kubernetes.
C'est un "TL;DR" ou une référence rapide d'informations utiles pour améliorer votre expérience de contribution sur GitHub.

**Table des matières**

- [Ressources utiles](#Ressources-utiles)
  - [Commencer](#Commencer)
  - [SIGs et autres groupes](#SIGs-et-autres-groupes)
  - [Communauté](#Communauté)
  - [Alias de messagerie importants](#Alias-de-messagerie-importants)
  - [Workflow](#Workflow)
  - [Tests](#Tests)
  - [Autres liens utiles](#Autres-liens-utiles)
- [Communiquer efficacement sur GitHub](#Communiquer-efficacement-sur-GitHub)
  - [Comment être excellent les uns envers les autres](#Comment-être-excellent-les-uns-envers-les-autres)
    - [Exemples de bonne mauvaise communication](#Exemples-de-bonne-mauvaise-communication)
- [Soumettre une contribution](#Soumettre-une-contribution)
  - [Signature de la CLA](#Signature-de-la-CLA)
  - [Ouverture et réponse aux Issues](#Ouverture-et-réponse-aux-Issues)
    - [Créer une Issue](#Créer-une-Issue)
    - [Répondre à une Issue](#Répondre-à-une-Issue)
  - [Ouverture d'une Pull Request](#Ouverture-d-une-Pull-Request)
    - [Créer une Pull Request](#Créer-une-Pull-Request)
    - [Exemple d'une description de Pull Request](#Exemple-d'une-description-de-Pull-Request)
    - [Dépannage d'une Pull Request](#Dépannage-d'une-Pull-Request)
  - [Labels](#Labels)
- [Travailler localement](#Travailler-localement)
  - [Stratégie de branche](#Stratégie-de-branche)
    - [Ajouter Upstream](#Ajouter-Upstream)
    - [Garder votre Fork synchronisé](#Garder-votre-Fork-synchronisé)
  - [Squashing Commits](#Squashing-Commits)

---

## Ressources utiles

### Commencer

- [Contributor Guide] - Guide sur la façon de commencer à contribuer au projet Kubernetes.
- [Developer Guide] - Guide pour contribuer du code directement au projet Kubernetes.

### SIGs et autres groupes

- [Liste principale des groupes][sigs]

### Communauté

- [Calendar] - Voir tous les événements de la communauté Kubernetes (réunions SIG / WG, événements, etc.)
- [kubernetes-dev] - La liste de diffusion sur le développement de Kubernetes
- [Kubernetes Forum] - Forum officiel de Kubernetes.
- [Slack channels] - Slack officiel de Kubernetes.
- [StackOverflow] - Un endroit pour poser vos questions d'utilisateur final de Kubernetes.
- [YouTube Channel] - Chaine officielle de la communauté Kubernetes.

### Workflow

- [Gubernator Dashboard] - Voir les Pull Requests entrantes et sortantes qui nécessitent votre attention.
- [Prow] - Kubernetes CI/CD System.
- [Tide] - Prow plugin that manages merges and tests. [Tide Dashboard]
- [Bot commands] - Commands used to interact with Kubernetes Bots (examples:
  `/cc`, `/lgtm`, and `/retest`)
- [GitHub labels] - Liste des labels utilisées dans le projet Kubernetes
- [Kubernetes Code Search], maintenu par [@dims]

### Tests

- [Prow] - Kubernetes CI/CD System.
- [Test Grid] - Afficher les tests historiques et leurs informations associées.
- [Triage Dashboard] - Regroupe les défaillances similaires pour un meilleur dépannage.
- [Velodrome] - Tableau de bord pour suivre le travail et tester la santé.

### Alias de messagerie importants

| Alias                          | Description                                                                                                                     |   |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------|---|
| community@kubernetes.io        | Envoyez un courrier électronique à l’équipe de la communauté (SIG Contributor Experience) au sujet d’un problème de communauté. |   |
| conduct@kubernetes.io          | Contactez le comité du code de conduite, liste de diffusion privée.                                                             |   |
| steering@kubernetes.io         | Postez le comité de pilotage. Adresse publique avec archive publique.                                                           |   |
| steering-private@kubernetes.io | Contacter le steering comité en privé, pour les sujets sensibles.                                                               |   |
| social@cncf.io                 | Contacter l'équipe sociale de la CNCF; blog, compte twitter et autres réseaux sociaux.                                          |   |

### Autres liens utiles

- [Statistiques de développeur] - Consultez les statistiques des développeurs pour tous les projets gérés par le CNCF.

---

## Communiquer efficacement sur GitHub

### Comment être excellent les uns envers les autres

Dans un premier temps, familiarisez-vous avec le [code de conduite].

#### Exemples de bonne / mauvaise communication

Quand on ouvre une issue, ou si vous avez besoin d’aide, soyez poli avec votre demande:

  🙂 "X ne compile pas quand je fais le Y, avez-vous des suggestions?"

  😞 «X ne marche pas! Réparez-ça, s'il vous plait!"

Lors de la fermeture d'une PR, transmettez un message explicatif et cordial expliquant pourquoi elle ne remplit pas les conditions requises pour être mergé.

🙂 «Je ferme ce PR car cette fonctionnalité ne peut pas prendre en charge le cas d’utilisation X. Dans le contexte proposé, il serait préférable de l’implémenter avec l’outil Y. Merci d'avoir travaillé sur cela. "

😞 «Pourquoi cela ne suit-il pas les conventions de l’API? Cela devrait être fait ailleurs!

---

## Soumettre une contribution

### Signature de la CLA

Avant de pouvoir soumettre une contribution, vous devez [signer le Contributor License Agreement(CLA)][cla].
Le projet Kubernetes ne peut accepter une contribution que si vous ou votre entreprise avez signé le CLA.

Si vous rencontrez des problèmes pour signer le CLA, suivez les [consignes de dépannage du CLA].

### Ouverture et réponse aux Issues

Les GitHub Issues sont le principal moyen de suivre des éléments tels que les rapports de bogues, les demandes d'amélioration ou de signaler d'autres problèmes tels que l'échec des tests.
Les issues ne sont **pas** destinées à être des [demandes de support utilisateur].
Pour ceux-ci, veuillez consulter le [guide de dépannage], signaler le problème à [stackOverflow] ou faire un suivi sur le [forum Kubernetes].

**References:**

- [Labels]
- [Prow commands][commands]

#### Créer un Issue

- Utilisez un Issuee template s'il en existe un. Utiliser le bon aidera d'autres contributeurs à répondre à votre problème.
  - Suivez les instructions décrites dans le template d'issue lui-même.
- Soyez descriptif avec la question que vous soulevez.
- Attribuer les [labels] appropriés. Si vous n'êtes pas sûr, le [k8s-ci-robot][prow] bot ([Kubernetes CI bot][prow]) répondra à votre problème avec les étiquettes nécessaires à son tri efficace.
- Soyez sélectif lorsque vous attribuez des Issues à l'aide de [`/assign @<username>`][assign] ou
  [`/cc @<username>`][cc]. Votre Issue sera triée plus efficacement en appliquant les labels corrects sur l'affectation de plus de personnes à la question.

#### Responding to an Issue

- Lorsque vous abordez un problème, laissez-le savoir aux autres sur lesquels vous travaillez cela pour éviter le travail en double.
- Lorsque vous avez résolu quelque chose par vous-même à un moment ultérieur, commentez la question de faire savoir aux gens avant de la fermer.
- Inclure des références à d’autres demandes PullRequests ou Issues (ou à tout matériel accessible),
  Exemple: _"ref: #1234"_. Il est utile d’identifier que des travaux connexes ont été résolu quelque part ailleurs.

### Ouverture d'une Pull Request

Les Pull requests (PR) sont les principaux moyens de contribuer au code, à la documentation ou à d’autres formes de travail qui seraient stockés dans un dépôt git.

**References:**

- [Labels]
- [Prow commands][commands]
- [Pull request process]
- [Github workflow]

#### Création d'une Pull Request

- Follow the directions of the pull request template if one is available. Cela aidera ceux qui répondent à votre PullRequest.
- Si un [correctif trivial] tel qu'un lien brisé, une faute de frappe ou une faute de grammaire, examinez l'ensemble du document pour rechercher d'autres erreurs potentielles. Ne pas ouvrir plusieurs PullRequests pour les petites corrections dans le même document.
- Référencez tous les problèmes liés à votre PullRequest ou les problèmes que PullRequest peut résoudre.
- Évitez de créer des modifications trop volumineuses dans un seul commit. Au lieu de cela, divisez votre PullRequest en plusieurs petits commits logiques. Cela facilite la révision de votre PullRequest.
- Commentez votre propre PullRequest lorsque vous pensez que quelque chose peut nécessiter une explication.
- Soyez sélectif lorsque vous affectez votre PullRequest avec [`/assign @<username>`][assign].
  L'affectation d'un nombre excessif de réviseurs ne donnera pas une révision plus rapide de PullRequest.
- Si votre PR est considéré comme un _"Work in progress"_ ajoutez un prefixe dans son nom avec `[WIP]` ou utilisez la commande [`/hold`][hold]. Ceci empêchera le merge de la PR jusqu'à la levée du `[WIP]` ou le retrait du hold.
- Si votre demande PullRequest n'a pas été relue, ne la fermez pas et n'ouvrez pas une nouvelle demande PullRequest avec les mêmes modifications. Notifiez les relecteurs dans un commentaire avec `@<github username>`.

#### Example PR Description

```text
Ref. #3064 #3097
All files owned by SIG testing were moved from `/devel` to the new folder `/devel/sig-testing`.

/sig contributor-experience
/cc @stakeholder1 @stakeholder2
/kind cleanup
/area developer-guide
/assign @approver1 @approver2 @approver3
```

Quel est le contenu de cette PR:

- **Line 1** - Référence à d'autres issues ou PRs (#3064 #3097).
- **Line 2** - Une brève description de ce qui se fait dans la PR.
- **Line 4** - Assignement au [SIG][sigs] avec la [commande][commands]
  `/sig contributor-experience`..
- **Line 5** - Les examinateurs qui peuvent avoir un intérêt sur cette issue ou PR sont spécifiés avec la commande [`/cc`][cc].
- **Line 6** - La commande [`/kind cleanup`][kind] ajoute un [label][labels] qui catégorise l'issue ou la PR en rapport avec le nettoyage du code, du processus ou de la dette technique.
- **Line 7** - La commande [`/area developer-guide`][kind] catégorise une issue ou PR en relation avec le guide du développeur.
- **Line 8** - La commande [`/assign`][assign] assigne un approbateur à la PR.
  Un approbateur sera suggéré par le [k8s-ci-robot][prow] est sélectionné dans la liste des propriétaires définis dans le fichier [OWNERS]. Ils vont ajouter le label [`/approve`][approve] à la PR après l'avoir passé en revue.

#### Troubleshooting a Pull Request

Après la proposition de votre PR, une série de tests est exécutée par la plateforme Kubernetes CI, [Prow].
Si l’un des tests échoue, le [k8s-ci-robot][prow] répondra à la PR avec des liens vers les tests ayant échoué et les journaux disponibles.

Pousser de nouveaux commits vers votre PR va automatiquement déclencher la ré-exécution des tests.

Il peut parfois y avoir des problèmes avec la plate-forme Kubernetes CI.
Celles-ci peuvent survenir pour diverses raisons même si votre contribution réussit tous les tests locaux.
Vous pouvez déclencher une nouvelle exécution des tests avec la commande `/retest`.

Pour plus d'informations sur le dépannage de tests spécifiques, voir le [Guide de test].

### Labels

Kubernetes utilise [étiquettes] pour catégoriser et trier les Issues et PullRequests.
L'application de labels appropriées aidera votre Issue ou PullRequest à être triée plus efficacement.

**References:**

- [Labels]
- [Prow commands][commands]

Labels fréquemment utilisés:

- [`/sig <sig name>`][kind] Attribuer un [SIG][SIGs] à la propriété de l'issue ou de la PR.
- [`/area <area name>`][kind] Associate the issue or PRs to a specific [area][labels].
- [`/kind <category>`][kind] [Categorizes][labels] the issue or PR.

---

## Travailler localement

Avant d'ouvrir une Pull Request, vous devrez effectuer préparer votre travail localement.
Si vous êtes nouveau sur git, le [tutoriel Atlassian git] est un bon point de départ.
En guise d'alternative, le didacticiel [Git magic] de Stanford est une bonne option multilingue.

**References:**

- [Atlassian git tutorial]
- [Git magic]
- [Github workflow]
- [Testing locally]
- [Developer guide]

### Stratégie de branche

Le projet Kubernetes utilise un workflow _"Fork and Pull"_ standard pour GitHub.
Dans le vocabulaire de git, votre fork personnel est appellée _"`origin`"_ et le dépôt git de référence du projet est appellé _"`upstream`"_.
Garder votre branche personnelle (`origin`) à jour avec le projet (`upstream`), il doit être configuré dans votre dépôt local.

#### Ajouter Upstream

Ajoutez `upstream` en tant que remote et configurez-le afin que vous ne puissiez pas y accéder.

```shell
# replace <upstream git repo> with the upstream repo url
# example:
#  https://github.com/kubernetes/kubernetes.git
#  git@github.com/kubernetes/kubernetes.git

git remote add upstream <upstream git repo>
git remote set-url --push upstream no_push
```

Cela peut être vérifié en exécutant `git remote -v` qui listera vos remotes configurées.

#### Garder votre dépôt synchronisé

Récupérez toutes les modifications de `upstream` et _"rebase"_ sur votre branche `master` locale.
Cela synchronisera votre dépôt local avec le projet `upstream`.

```text
git fetch upstream
git checkout master
git rebase upstream/master
```

Effectuez cette opération au minimum avant de créer une nouvelle branche pour travailler sur votre fonctionnalité ou votre correctif.

```text
git checkout -b myfeature
```

#### Squashing Commits

Le but principal de [squashing commits] est de créer un historique git lisible.
Cela se fait généralement dans la dernière phase d'une PullRequest.
Si vous ne savez pas si vous devez faire un squash de vos commits, il est préférable de préférer avoir plus de commits et de laisser le soin aux autres contributeurs de réviser et d’approuver vos PullRequests.

[guide du contributeur]: /contributors/guide/README.md
[guide du développeur]: /contributors/devel/README.md
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
[statistiques de développeur]: https://k8s.devstats.cncf.io
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
[Guide de test]: /contributors/devel/sig-testing/testing.md
[labels]: https://git.k8s.io/test-infra/label_sync/labels.md
[solution triviale]: /contributors/guide/pull-requests.md#10-trivial-edits
[Github workflow]: /contributors/guide/github-workflow.md#3-branch
[squashing commits]: /contributors/guide/pull-requests.md#6-squashing-and-commit-titles
[owners]: /contributors/guide/owners.md
[tester localement]: /contributors/guide/README.md#testing
[developer guide]: /contributors/devel/README.md
[Atlassian git tutorial]: https://www.atlassian.com/git/tutorials
[git magic]: http://www-cs-students.stanford.edu/~blynn/gitmagic/
