# Kubernetesコントリビューターチートシート

Kubernetesにコントリビュートする際のtipsや、Kubernetesプロジェクト内で使用されているベストプラクティスなどの共通リソースのリストです。
これらのまとめや便利な情報へのクイックリファレンスはGitHubでのコントリビューションの体験をよりよいものにすることでしょう。

**目次**
- [便利なリソース](#便利なリソース)
  - [はじめに](#はじめに)
  - [SIGとその他のグループ](#SIGとその他のグループ)
  - [コミュニティ](#コミュニティ)
  - [重要なEメールエイリアス](#重要なEメールエイリアス)
  - [ワークフロー](#ワークフロー)
  - [テスト](#テスト)
  - [その他の便利なリンク](#その他の便利なリンク)
- [GitHub上での効率的なコミュニケーション](#GitHub上での効率的なコミュニケーション)
  - [お互いに良くあるためにはどうしたらよいか](#お互いに良くあるためにはどうしたらよいか)
    - [良いまたは悪いコミュニケーションの例](#良いまたは悪いコミュニケーションの例)
- [貢献する](#貢献する)
  - [CLAにサインする](#CLAにサインする)
  - [Issueを開いたり返事をしたりする](#Issueを開いたり返事をしたりする)
    - [Issueを作る](#Issueを作る)
    - [Issueに返事をする](#Issueに返事をする)
  - [Pull Requestを開く](#Pull-Requestを開く)
    - [Pull Requestを作成する](#Pull-Requestを作成する)
    - [PRの説明文の例](#PRの説明文の例)
    - [Pull Requestのトラブルシューティング](#Pull-Requestのトラブルシューティング)
  - [ラベル](#ラベル)
- [ローカルでの作業](#ローカルでの作業)
  - [ブランチ戦略](#ブランチ戦略)
    - [Upstreamを追加する](#Upstreamを追加する)
    - [フォークを最新に保つ](#フォークを最新に保つ)
  - [コミットをまとめる](#コミットをまとめる)

---

## 便利なリソース

### はじめに

- [コントリビューターガイド] - Kubernetesプロジェクトへコントリビュートする方法のガイド
- [開発者ガイド] - Kubernetesプロジェクトへコードを直接コントリビュートする方法のガイド
- [セキュリティと情報開示] - 脆弱性の報告とセキュリティリリースプロセスのガイド

### SIGとその他のグループ

- [グループのリスト][sigs]

### コミュニティ

- [カレンダー] - Kubernetesコミュニティでのイベントの一覧(SIG/WGのミーティングやイベントなど)
- [kubernetes-dev] - Kubernetes開発メーリングリスト
- [Kubernetesフォーラム] - Kubernetesの公式フォーラム
- [Slackチャンネル] - Kubernetesの公式Slack
- [Stack Overflow] - Kubernetesのエンドユーザーとしての質問を聞く場所
- [YouTubeチャンネル] - Kubernetesコミュニティの公式チャンネル


### ワークフロー

- [Gubernatorダッシュボード] - 注意して見ておくべきPull Requests
- [Prow] - KubernetesのCI/CDシステム
- [Tide] - mergeやtestを管理するためのProw用プラグイン [Tideダッシュボード]
- [Botコマンド] - KubernetesのBotとコミュニケーションをとるためのコマンド (例: `/cc`、`/lgtm`や`/retest`)
- [GitHubラベル] - Kubernetesプロジェクトで使用されるラベルのリスト
- [@dims]によって保守されている[Kubernetes Code Search]


### テスト

- [Prow] - KubernetesのCI/CDシステム
- [Test Grid] - 歴史的なテストや関連した情報を見る
- [Triageダッシュボード] - よりよくトラブルシューティングをするために、似たような失敗をまとめる
- [Velodrome] - ジョブやテスト結果を追跡するためのダッシュボード


### 重要なEメールエイリアス

- community@kubernetes.io - コミュニティの問題について、コミュニティチーム(SIG Contributor Experience)の誰かにメールするアドレス
- conduct@kubernetes.io - 行動規範委員会へ連絡を取るためのプライベートメーリングリスト
- steering@kubernetes.io - 運営委員会へメールするアドレスで、公開アーカイブのある公開アドレス
- steering-private@kubernetes.io - 運営委員会へセンシティブなことを伝えるためのプライベートアドレス
- social@cncf.io - CNCFソーシャルチームへの連絡先(blogやtwitterアカウントなど)


### その他の便利なリンク

- [開発者統計] - CNCFが管理するプロジェクトの開発者統計情報

---

## GitHub上での効率的なコミュニケーション


### お互いに良くあるためにはどうしたらよいか

まず最初に[Code of Conduct]をよく読んでください。


#### 良いまたは悪いコミュニケーションの例

issueをあげる時や助けを求める時、礼儀正しくしてください:

  🙂 「Yをやった時にXがコンパイルできませんでした。なにかいい方法はありませんか？」

  😞 「Xが動かない！直して！」

PRを閉じるとき、どうしてmergeできないのか、誠心誠意説明し、伝えてください。

🙂 「この機能はXというユースケースをサポートしていないのでこのPRを閉じます。提案された形であれば、Yツールで実装される方がよりよいと思います。」

😞 「どうしてAPI規約に従っていないんですか？これは他でやるべきです！」

---

## 貢献する

### CLAにサインする

コントリビューションを提出する前に、[Contributor License Agreement(CLA)にサインする](cla)必要があります。Kubernetesプロジェクトは、あなたもしくはあなたの会社がCLAにサイン済みの場合にのみコントリビューションの受け入れを行います。

CLAのサインで何か問題があった場合、[CLAトラブルシューティングガイドライン]を参照してください。


### Issueを開いたり返事をしたりする

GitHub Issueはバグレポートや改善要求、あるいはテスト失敗のようなその他の問題を追跡するための最初の手段です。[ユーザーによるサポート要求]の方法としては使用**されていません**。そのような場合は[トラブルシューティングガイド]をみて、[Stack Overflow]や[Kubernetesフォーラム]に問題を報告してください。

**参考:**
- [ラベル]
- [Prowコマンド][コマンド]


#### Issueを作る

- もし用意されているなら、Issue templateを使用してください。適切なテンプレートを使用することで、他のコントリビューターが返信しやすくなります。
  - Issue template自体に書かれている手順に従ってください。
- 詳細な説明をIssueに記述してください。
- 適切な[ラベル]を設定してください。よくわからなければ、[k8s-ci-robot][prow]([Kubernetes CI bot][prow])というボットが、重要度を適切に判断するために必要なラベルを提案します。
- [`/assign @<username>`][assign]か[`/cc @<username>`][cc]を使用して担当者をアサインする場合は選択的に行ってください。より多くの人にアサインをするより、適切なラベルを付ける方が効果的です。


#### Issueに返事をする

- Issueに取り組む時は、他の人とバッティングしないように、コメントを残してください。
- 自己解決した場合には、Issueを閉じる前に他の人にわかるようコメントしてください。
- 他のPRやIssue(あるいはその他アクセス可能なもの)への参照を含めてください(例えば、 _"ref: #1234"_ のように)。他の場所にある関連した作業を特定するのに便利です。


### Pull Requestを開く

Pull request(PR)はコード、ドキュメント、あるいはgitリポジトリに格納されているその他のものに対してコントリビュートする際の主な手段です。

**参考:**
- [ラベル]
- [Prowコマンド][コマンド]
- [Pull request process]
- [GitHub workflow]


#### Pull Requestを作成する

- 利用可能な場合、Pull Requestテンプレートの指示に従います。 それはあなたのPRに対応する人々の助けになります。
- リンク切れやタイプミス、文法の間違いなどの[簡単な修正]の場合、他の可能性のある間違いについてドキュメント全体を見直してください。
  同じドキュメントの小さな修正で複数のPRを作成しないでください。
- PRに関連するIssueやPRで解決する可能性があるIssueを参照してください。
- 一度のコミットで過大な変更を加えないでください。代わりに、PRを複数の小さなコミットに分割してください。
  これによりPRのレビューが容易になります。
- 何か説明を加える必要があると思われる場合は、PRにコメントしてください。
- [`/assign @<username>`][assign]でPRに割り当てるときは選択的にしてください。
  過剰なレビュー担当者を割り当てたからといって、 迅速なレビューが得られるわけではありません。
- あなたのPRが _"進行中"_ とされる場合、名前の前に `[WIP]` を付けるか、[`/hold`][hold]コマンドを使用してください。これは `[WIP]` またはHoldが解除されるまでPRがマージされるのを防ぎます。
- あなたのPRがレビューされてない場合に、閉じて同じ変更のPRを新しく作成しないでください。`@<github username>` とコメントでレビュアーにPingしてください。



#### PRの説明文の例

```
Ref. #3064 #3097
All files owned by SIG testing were moved from `/devel` to the new folder `/devel/sig-testing`.

/sig contributor-experience
/cc @stakeholder1 @stakeholder2
/kind cleanup
/area developer-guide
/assign @approver1 @approver2 @approver3
```

PRの内容:
- **1行目** - 他のIssueやPRへの参照(#3064 #3097)
- **2行目** - PRで行われていることの簡単な説明
- **4行目** - `/sig contributor-experience` [コマンド]での[SIG][sigs]の割り当て
- **5行目** - この特定のIssueやPRに関心があるレビュアーを[`/cc`][cc]コマンドで指定
- **6行目** - [`/kind cleanup`][kind]コマンドでコードやプロセス、技術的負債の整理に関してIssueやPRを分類する[ラベル][ラベル]を追加
- **7行目** - [`/area developer-guide`][kind]コマンドで開発者ガイドに関してIssueやPRを分類
- **8行目** - [`/assign`][assign]コマンドでPRにApproverを割り当て。
  Approverは[k8s-ci-robot][prow]によって提案され、[OWNERS]ファイルのオーナーのリストから選択されます。
  Approverはレビューされた後のPRに[`/approve`][approve]ラベルを追加します

#### Pull Requestのトラブルシューティング

PRが作成された後、KubernetesのCIプラットフォームの[Prow]によって一連のテスト
が実行されます。テストのいずれかが失敗した場合、[k8s-ci-robot][prow]は
失敗したテストへのリンクと有効なログをPRに返信します。

新しいコミットをPRにがプッシュすると、自動的にテストが再実行されます。

時折KubernetesのCIプラットフォームに問題がある場合があります。
あなたの貢献が全てのローカルテストに合格したとしても、
これらは様々な理由で発生する場合があります。`/retest` コマンドでテストを
再実行することができます。

特定のテストのトラブルシューティングの詳細については[テストガイド]を参照してください。

### ラベル

KubernetesはIssueとPull Requestを分類し、優先順位を付けるために[ラベル]を使用します。
正しいラベルを付けることであなたのIssueやPRをより効果的に処理することができます。


**参考:**
- [ラベル]
- [Prowコマンド][コマンド]

よく使われるラベル:
- [`/sig <sig name>`][kind]はIssueやPRのアサインを[SIG][SIGs]に割り当てます。
- [`/area <area name>`][kind]はIssueやPRを特定の[分野][ラベル]に関連付けます。
- [`/kind <category>`][kind]はIssueやPRを[分類][ラベル]します。

---

## ローカルでの作業

Pull Requestを作成する前に、ローカルである程度の作業を行う必要があります。
もしあなたがgitに慣れていない場合、[Atlassian gitチュートリアル]は良い出発点です。
他に、Stanfordの[Git magic]チュートリアルは多言語に対応しています。

**参考:**
- [Atlassian gitチュートリアル]
- [Git magic]
- [GitHub workflow]
- [ローカルでのテスト]
- [開発者ガイド]


### ブランチ戦略

KubernetesプロジェクトはGitHubの標準である _"Fork and Pull"_ ワークフローを使います。
gitの用語で、あなたの個人的なフォークは _"`origin`"_ と呼ばれ、実際のプロジェクトの
gitリポジトリのことを _"`upstream`"_ と呼ばれます。
あなたの個人的なブランチ(`origin`)をプロジェクト(`upstream`)の最新に保つために、
ローカルの作業コピー内で設定されてなければなりません。


#### Upstreamを追加する

`upstream` をリモートとして追加し、プッシュできないように設定してください。

```
# replace <upstream git repo> with the upstream repo url
# example:
#  https://github.com/kubernetes/kubernetes.git
#  git@github.com/kubernetes/kubernetes.git

git remote add upstream <upstream git repo>
git remote set-url --push upstream no_push
```

これは設定したリモートを一覧表示する `git remote -v` を実行して確認することができます。


#### フォークを最新に保つ

`upstream` から全ての変更を取得し、ローカルの `master` ブランチに _"rebase"_ します。
これはローカルのリポジトリを `upstream` プロジェクトと同期させます。

```
git fetch upstream
git checkout master
git rebase upstream/master
```

あなたは機能への取り組みや修正をするためにブランチを作成する前に、最低限これをやるべきです。

```
git checkout -b myfeature
```

#### コミットをまとめる

[スカッシュコミット]の主な目的はきれいに読めるgitの履歴や加えられた変更のログを作成することです。通常これはPRの改定の最終段階に行われます。
コミットをスカッシュするべきかわからない場合は、作業を止めてPRのレビューと承認を担当する他のコントリビューターの判断に任せることをおすすめします。


[コントリビューターガイド]: /contributors/guide/README.md
[開発者ガイド]: /contributors/devel/README.md
[gubernatorダッシュボード]: https://gubernator.k8s.io/pr
[prow]: https://prow.k8s.io
[tide]: http://git.k8s.io/test-infra/prow/cmd/tide/pr-authors.md
[tideダッシュボード]: https://prow.k8s.io/tide
[botコマンド]: https://go.k8s.io/bot-commands
[GitHubラベル]: https://go.k8s.io/github-labels
[Kubernetes Code Search]: https://cs.k8s.io/
[@dims]: https://github.com/dims
[カレンダー]: https://calendar.google.com/calendar/embed?src=cgnt364vd8s86hr2phapfjc6uk%40group.calendar.google.com
[kubernetes-dev]: https://groups.google.com/forum/#!forum/kubernetes-dev
[slackチャンネル]: http://slack.k8s.io/
[Stack Overflow]: https://stackoverflow.com/questions/tagged/kubernetes
[youtubeチャンネル]: https://www.youtube.com/c/KubernetesCommunity/
[triageダッシュボード]: https://go.k8s.io/triage
[test grid]: https://testgrid.k8s.io
[velodrome]: https://go.k8s.io/test-health
[開発者統計]: https://k8s.devstats.cncf.io
[code of conduct]: /code-of-conduct.md
[ユーザーによるサポート要求]: /contributors/guide/issue-triage.md#determine-if-its-a-support-request
[トラブルシューティングガイド]: https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/
[kubernetesフォーラム]: https://discuss.kubernetes.io/
[pull request process]: /contributors/guide/pull-requests.md
[github workflow]: /contributors/guide/github-workflow.md
[prow]: https://git.k8s.io/test-infra/prow#prow
[cla]: /CLA.md#how-do-i-sign
[claトラブルシューティングガイドライン]: /CLA.md#troubleshooting
[コマンド]: https://prow.k8s.io/command-help
[kind]: https://prow.k8s.io/command-help#kind
[cc]: https://prow.k8s.io/command-help#cc
[hold]: https://prow.k8s.io/command-help#hold
[assign]: https://prow.k8s.io/command-help#assign
[SIGs]: /sig-list.md
[テストガイド]: /contributors/devel/sig-testing/testing.md
[ラベル]: https://git.k8s.io/test-infra/label_sync/labels.md
[簡単な修正]: /contributors/guide/pull-requests.md#10-trivial-edits
[GitHub workflow]: /contributors/guide/github-workflow.md#3-branch
[スカッシュコミット]: /contributors/guide/pull-requests.md#6-squashing-and-commit-titles
[owners]: /contributors/guide/owners.md
[ローカルでのテスト]: /contributors/guide/README.md#testing
[Atlassian gitチュートリアル]: https://www.atlassian.com/git/tutorials
[git magic]: http://www-cs-students.stanford.edu/~blynn/gitmagic/
[セキュリティと情報開示]: https://kubernetes.io/docs/reference/issues-security/security/
[approve]: https://prow.k8s.io/command-help#approve
