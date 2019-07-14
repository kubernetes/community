# Cheat Sheet do Colaborador do Kubernetes

Uma lista de recursos comuns ao contribuir para o Kubernetes, dicas, truques e
melhores melhores práticas comumente usadas no projeto Kubernetes. É um resumo ou
referência rápida de informações úteis para tornar a sua experiência de contribuição do GitHub
Melhor.

**Índice**
- [Recursos úteis](#recursos-úteis)
  - [Primeiros passos](#primeiro-passos)
  - [SIGs e outros grupos](#sigs-e-outros-grupos)
  - [Comunidade](#comunidade)
  - [E-mails Importantes](#e-mails-importantes)
  - [Fluxo de trabalho](#fluxo-de-trabalho)
  - [Testes](#testes)
  - [Outros Links Úteis](#outros-links-úteis)
- [Comunicar efetivamente no GitHub](#comunicando-efetivamente-no-github)
  - [Como ser excelente um para o outro](#como-ser-excelente-para-outro)
    - [Exemplos de Comunicação Boa/Má](#exemplos-de-boa-e-ma-comunicação)
- [Enviando uma contribuição](#enviando-uma-contribuição)
  - [Assinando o CLA](#assinando-o-cla)
  - [Abrindo e respondendo a issues](#abrindo-e-respondendo-a-issues)
    - [Criando uma Issue](#criando-uma-issue)
    - [Respondendo a uma Issue](#respondendo-a-uma-issue)
  - [Abrindo uma Solicitação de Pull Request](#abrindo-um-pull-request)
    - [Criando um Pull Request](#criando-um-pull-request)
    - [Exemplo de descrição Pull Request](#exemplo-de-descrição-de-pull-request)
    - [Solucionando problemas de Pull Request](#solucionando-problemas-de-pull-request)
  - [Labels](#labels)
- [Trabalhando localmente](#trabalhando-localmente)
  - [Estratégia de Branch](#estratégia-de-branch)
    - [Adicionando Upstream](#adicionando-upstream)
    - [Mantendo seu fork em sincronia](#mantendo-seu-fork-em-sincronia)
  - [Agrupando os Commits (squash)](#agrupando-os-commits)
---

## Recursos úteis

### Primeiro passos

- [Guia do Contribuidor] - Guia sobre como começar a contribuir para o projeto Kubernetes.
- [Guia do Desenvolvedor] - Guia para contribuir com código diretamente para o projeto Kubernetes.
- [Informações de Segurança e Divulgações] - Guia para relatar vulnerabilidades e o processo de release.

### SIGs e Outros Grupos

- [Lista de grupos Master][sigs]

### Comunidade

- [Calendário] - Ver todos os eventos da Comunidade Kubernetes (reuniões SIG / WG,
  eventos etc.)
- [kubernetes-dev] - A lista de discussão do desenvolvimento do Kubernetes
- [Fórum do Kubernetes] - Fórum oficial do Kubernetes.
- [Slack] - Slack Oficial do Kubernetes.
- [Stack Overflow] - Um lugar para fazer perguntas ao usuário final do Kubernetes.
- [YouTube] - Canal oficial da comunidade Kubernetes.


### Fluxo de trabalho

- [Gubernator Dashboard] - Ver Pull Requests que exigem sua atenção.
- [Prow] - Kubernetes CI/CD System.
- [Tide] - Plugin Prow que gerencia merges e testes. [Tide Dashboard]
- [Comandos do Bot] - Comandos usados ​​para interagir com o Kubernetes Bots (exemplos:
  `/cc`, `/lgtm`, and `/retest`)
- [GitHub labels] - Lista de lebels usados ​​em todo o projeto Kubernetes
- [Pesquisa no código do Kubernetes], mantido por [@dims]


### Testes

- [Prow] - Kubernetes CI/CD System.
- [Test Grid] - Veja o histórico de testes e suas informações associadas.
- [Dashboard de Triagem] - Junta falhas semelhantes para melhor solução de problemas.
- [Velodrome] - Dashboard para rastrear jobs e testar a estabilidade.


### E-mails-Importantes

- community@kubernetes.io - Envie uma mensagem para alguém da equipe (Colaborador com experiência na SIG) sobre algum problema da comunidade.
- conduct@kubernetes.io - Entre em contato com o comitê do Código de Conduta, atravéz do mailing list privado.
- steering@kubernetes.io - Envie uma mensagem para o comitê diretor. Endereço público com arquivo público.
- steering-private@kubernetes.io - Envie mensagem para o comitê diretor privativo, para itens sensíveis.
- social@cncf.io - Entre em contato com a equipe social da CNCF: blog, conta do twitter e
  outras propriedades sociais.


### Outros Links Úteis

- [Estatísticas do Desenvolvedor] - Veja as estatísticas do desenvolvedor para todos os projetos da CNCF gerenciados.

---

## Comunicando Efetivamente no GitHub


### Como ser excelente para outro

Como primeiro passo, familiarize-se com o [Código de Conduta].


#### Exemplos de boa e ma Comunicação

Ao levantar um problema ou solicitar assistência, seja educado com sua solicitação:

  🙂 “X não compila quando eu faço Y, você tem alguma sugestão?”

  😞 “X não funciona! Por favor conserte!”

Ao fechar um PR, transmita uma mensagem explicativa e cordial explicando
por que não atende aos requisitos a serem mesclados.

🙂 “Estou fechando este PR porque esse recurso não suporta o caso de uso X. Seria melhor ser implementado com a ferramenta Y. Obrigado você por trabalhar nisso.”

😞 “Por que isso não segue as convenções da API? Isso deve ser feito em outro lugar!”

---

## Enviando uma contribuição

### Assinando o CLA

Antes de enviar uma contribuição, você deve [assinar o Contributor License
Agreement(CLA)][cla]. O projeto Kubernetes só pode aceitar uma contribuição
se você ou sua empresa assinou o CLA.

Se você encontrar algum problema ao assinar o CLA, veja o [solucionando problemas do cla].


### Abrindo e respondendo a issues

GitHub Issues é o principal meio de rastrear coisas como relatórios de bugs,
Pull Requests ou relatar outros problemas (issues), como testes com falha. Eles
**não** são destinados a [solicitações de suporte ao usuário]. Para suporte, por favor, verifique com o
[guia de solução de problemas], relate o problema para o [Stack Overflow] ou faça o acompanhamento
no [Fórum do Kubernetes].

**References:**
- [Labels]
- [Prow commands][commands]


#### Criando uma issue

- Use um template de uma issue, se houver algum disponível. Usar corretamente ajudará outros
  contribuidores para responder a sua issue.
  - Siga as instruções descritas no próprio modelo de assunto.
- Seja descritivo com a issue que você está criando.
- Atribuir [labels] apropriadas. Se você não tiver certeza, o [k8s-ci-robot][prow] bot
  ([Kubernetes CI bot][prow]) responderá ao seu problema com os rótulos necessários
  para que seja realizado uma triagem efetiva.
- Seja seletivo ao atribuir problemas usando[`/assign @<username>`][assign] ou
  [`/cc @<username>`][cc]. Sua issue passará por uma triagem mais efetiva se utilizar as labels 
  e atribuir a issue a mais pessoas.


#### Respondendo a uma issue

- Ao lidar com uma issue, comente sobre ela e deixe que outros saibam que você está trabalhando nela
  para evitar trabalho duplicado.
- Quando você resolver algo por conta própria em qualquer momento futuro, comente
  a issue deixando as pessoas saberem antes de fechá-lo.
- Inclua referências a outros PRs ou questões (ou quaisquer materiais acessíveis),
  exemplo: _"ref: #1234"_. É útil identificar que o trabalho relacionado foi
  endereçado em outro lugar.


### Abrindo um Pull Request

Pull Request (PR) é o principal meio de contribuir com código, documentação ou
outras formas de trabalho que seriam armazenadas em um repositório git.

**References:**
- [Labels]
- [Prow commands][commands]
- [Processo de pull request]
- [GitHub workflow]


#### Criando um Pull Request

- Siga as instruções do PR, se houver um disponível. Isto
  vai ajudar aqueles que respondem ao seu PR.
- Se uma [correção trivial], como erro de link, erro ortográfico ou gramática, revise o
  documento inteiro para outros possíveis erros. Não abra vários PRs para
  pequenas correções no mesmo documento.
- Faça referência a quaisquer problemas relacionados ao seu PR ou a problemas que o PR possa resolver.
- Evite criar alterações excessivamente grandes em um único commit. Em vez disso, interrompa seu PR
  em vários pequenos pedaços lógicos. Isso torna mais fácil para o seu PR ser revisado.
- Comente sobre o seu próprio PR onde você acredita que algo pode precisar de mais explicação.
- Seja seletivo ao atribuir seu PR com[`/assign @<username>`][assign].
  Atribuir revisores excessivos não resultará em uma revisão de RP mais rápida.
- Se o seu PR é considerado _"Work in progress"_ com o prefixo`[WIP]` 
ou use o command[`/hold`][hold]. Isso impedirá que o PR seja mergeado
  até o `[WIP]` ser retirado.
- Se você não teve seu PR revisado, não feche e abra um novo PR com o
 mesmas mudanças. Ping seus revisores em um comentário com `@<github username>`.


#### Exemplo de descrição de Pull Request

```
Ref. #3064 #3097
Todos os arquivos pertencentes ao teste SIG foram movidos de `/devel` para a nova pasta `/devel/sig-testing`.

/sig contributor-experience
/cc @stakeholder1 @stakeholder2
/kind cleanup
/area developer-guide
/assign @approver1 @approver2 @approver3
```

O que está nesse PR:
- **Line 1** - Referência a outras questões ou PRs (#3064 #3097).
- **Line 2** - Uma breve descrição do que está sendo feito no PR.
- **Line 4** - [SIG][sigs] atribuição com o [commando][commands]
  `/sig contributor-experience`..
- **Line 5** - Revisores que possam ter interesse sobre esta questão específica ou PR são
  especificado com o comando [`/cc`][cc].
- **Line 6** - O comando [`/kind cleanup`][kind] adiciona um [label][labels] que
  categoriza a issue ou PR relacionado com a limpeza do código, processo ou
débito técnico.
- **Line 7** - O comando [`/area developer-guide`][kind] categoriza issues ou
  PRs relacionados com o guia do desenvolvedor.
- **Line 8** - O comando [`/assign`][assign] atribui um aprovador ao PR.
  Um aprovador será sugerido pelo [k8s-ci-robot][prow] e é selecionado de
  a lista de proprietários no arquivo [OWNERS]. Eles vão adicionar um label
  [`/approve`][approve] para o PR depois de ter sido revisto.


#### Solucionando problemas de Pull Request

Após o seu PR ser proposto, uma série de testes serão executados" pelo CI da plataforma Kubernetes,[Prow]. 
Se algum dos testes falhou, o [k8s-ci-robot][prow] responderá ao PR com links para os testes com falha e logs disponíveis.

Enviar novos commits para o seu PR irá disparar automaticamente os testes para serem executados novamente.

Ocasionalmente, pode haver problemas com o CI da plataforma Kubernetes. Estes podem ocorrer
por uma ampla variedade de razões, mesmo que sua contribuição passe por todos os
testes. Você pode acionar uma nova execução dos testes com o comando `/retest`.

Para obter mais informações sobre como solucionar problemas específicos, consulte o [guia para testes].


### Labels

O Kubernetes usa [labels] para categorizar e realizar uma triagem de issues e PRs.
A aplicação das labels corretas ajudará sua issue ou PR passar pela triagem
efetivamente.

**References:**
- [Labels]
- [Prow commands][commands]

Labels usados ​​com frequência:
- [`/sig <sig name>`][kind] Atribui a [SIG][SIGs] para a posse da issue
  ou PR.
- [`/area <area name>`][kind] Associe a issue ou PRs a uma
  [area][labels] específica.
- [`/kind <category>`][kind] [Categorizes][labels] A issue ou PR.

---

## Trabalhando Localmente

Antes de propor um PR, você terá que preparar seu ambiente local.
Se você é novo no git, o [Tutorial git Atlassian] é um bom começo.
Como alternativa, o tutorial de [Tutorial git magic] da Stanford é uma boa
opção multi-idioma.

**Referências:**
- [Tutorial git Atlassian]
- [Tutorial git magic]
- [GitHub workflow]
- [Testando localmente]
- [Guia do Desenvolvedor]


### Estratégia de Branch

O projeto Kubernetes usa um fluxo _"Fork and Pull"_ que é o padrão para o
GitHub. Em termos gerais, seu fork pessoal é chamado de _"`origin`"_ e
o repositório git do projeto é chamado _"`upstream`"_. Para manter seu
fork pessoal (`origin`) atualizado com o projeto (`upstream`), você deve
configurá-lo localmente.


#### Adicionando Upstream

Adicione o `upstream` como um repositório remoto e configure-o para que você não possa efetuar o git push para ele.

```
# substituir <upstream git repo> com a URL do repositório upstream
# exemplo:
#  https://github.com/kubernetes/kubernetes.git
#  git@github.com/kubernetes/kubernetes.git

git remote add upstream <upstream git repo>
git remote set-url --push upstream no_push
```

Isto pode ser verificado executando o comando `git remote -v` que listará seus
repositórios remotos.


#### Mantendo seu fork em sincronia

Busque todas as mudanças de `upstream` e faça o _"rebase"_ em sua branch `master` local.
Isso irá sincronizar seu repositório local com o projeto `upstream`.

```
git fetch upstream
git checkout master
git rebase upstream/master
```

Este é o mínimo que você deveria fazer antes de criar uma nova branch para trabalhar em sua
feature ou issue.

```
git checkout -b myfeature
```

#### Agrupando os commits

O objetivo principal de [agrupar os commits] (squashing) é criar um git limpo com
histórico legível e com log das alterações feitas. Geralmente isso é feito na última
fase de uma revisão do PR. Se você não tem certeza se deve efetuar o squashing em seus commits,
é melhor errar não efetuando o squashing, e deixar para o julgamento dos outros contribuidores designados à revisar e aprovar o seu PR.



[Guia do Contribuidor]: /contributors/guide/README.md
[Guia do Desenvolvedor]: /contributors/devel/README.md
[Gubernator dashboard]: https://gubernator.k8s.io/pr
[prow]: https://prow.k8s.io
[tide]: http://git.k8s.io/test-infra/prow/cmd/tide/pr-authors.md
[tide dashboard]: https://prow.k8s.io/tide
[Comandos do Bot]: https://go.k8s.io/bot-commands
[gitHub labels]: https://go.k8s.io/github-labels
[Pesquisa no código do Kubernetes]: https://cs.k8s.io/
[@dims]: https://github.com/dims
[Calendário]: https://calendar.google.com/calendar/embed?src=cgnt364vd8s86hr2phapfjc6uk%40group.calendar.google.com
[kubernetes-dev]: https://groups.google.com/forum/#!forum/kubernetes-dev
[slack]: http://slack.k8s.io/
[Stack Overflow]: https://stackoverflow.com/questions/tagged/kubernetes
[youtube]: https://www.youtube.com/c/KubernetesCommunity/
[Dashboard de Triagem]: https://go.k8s.io/triage
[test grid]: https://testgrid.k8s.io
[velodrome]: https://go.k8s.io/test-health
[Estatísticas do Desenvolvedor]: https://k8s.devstats.cncf.io
[Código de Conduta]: /code-of-conduct.md
[solicitações de suporte ao usuário]: /contributors/guide/issue-triage.md#determine-if-its-a-support-request
[guia de solução de problemas]: https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/
[Fórum do Kubernetes]: https://discuss.kubernetes.io/
[Processo de pull request]: /contributors/guide/pull-requests.md
[github workflow]: /contributors/guide/github-workflow.md
[prow]: https://git.k8s.io/test-infra/prow#prow
[cla]: /CLA.md#how-do-i-sign
[solucionando problemas do cla]: /CLA.md#troubleshooting
[commands]: https://prow.k8s.io/command-help
[kind]: https://prow.k8s.io/command-help#kind
[cc]: https://prow.k8s.io/command-help#cc
[hold]: https://prow.k8s.io/command-help#hold
[assign]: https://prow.k8s.io/command-help#assign
[SIGs]: /sig-list.md
[guia para testes]: /contributors/devel/sig-testing/testing.md
[labels]: https://git.k8s.io/test-infra/label_sync/labels.md
[trivial fix]: /contributors/guide/pull-requests.md#10-trivial-edits
[GitHub workflow]: /contributors/guide/github-workflow.md#3-branch
[agrupar os commits]: /contributors/guide/pull-requests.md#6-squashing-and-commit-titles
[owners]: /contributors/guide/owners.md
[Testando localmente]: /contributors/guide/README.md#testing
[Tutorial git Atlassian]: https://www.atlassian.com/git/tutorials
[Tutorial git magic]: http://www-cs-students.stanford.edu/~blynn/gitmagic/
[Informações de Segurança e Divulgações]: https://kubernetes.io/docs/reference/issues-security/security/
[approve]: https://prow.k8s.io/command-help#approve
