# 쿠버네티스 컨트리뷰터 참고 자료(Cheat Sheet)

쿠버네티스에 기여할 때 일반적인 자원, 팁, 요령과
쿠버네티스 프로젝트 내에서 사용되는 일반적인 모범 사례의 목록이다.
이 문서는 GitHub 컨트리뷰션 경험을 더 좋게 만들기 위한 "TL;DR" 혹은
유용한 정보의 빠른 레퍼런스이다.

**목차**
- [도움이 되는 자원](#도움이-되는-자원)
  - [시작하기](#시작하기)
  - [SIG 및 기타 그룹](#SIG-및-기타-그룹)
  - [커뮤니티](#커뮤니티)
  - [중요한 이메일 별칭](#중요한-이메일-별칭)
  - [작업 흐름](#작업-흐름)
  - [테스트](#테스트)
  - [다른 유용한 링크](#다른-유용한-링크)
- [GitHub에서 효과적으로 의사 소통하기](#GitHub에서-효과적으로-의사-소통하기)
  - [서로에게 훌륭한 방법](#서로에게-훌륭한-방법)
    - [좋은/나쁜 의사 소통 예시](#좋은/나쁜-의사-소통-예시)
- [컨트리뷰션 전송](#컨트리뷰션-전송)
  - [CLA 서명](#CLA-서명)
  - [이슈의 오픈 및 응답](#이슈의-오픈-및-응답)
    - [이슈 생성](#이슈-생성)
    - [이슈 응답](#이슈-응답)
  - [풀 리퀘스트 오픈](#풀-리퀘스트-오픈)
    - [풀 리퀘스트 생성](#풀-리퀘스트-생성)
    - [PR 설명 예시](#PR-설명-예시)
    - [풀 리퀘스트 문제 해결](#풀-리퀘스트-문제-해결)
  - [레이블](#레이블)
- [로컬에서 작업](#로컬에서-작업)
  - [브랜치 전략](#브랜치-전략)
    - [Upstream 추가](#Upstream-추가)
    - [Fork를 동기화 상태로 유지하기](#Fork를-동기화-상태로-유지하기)
  - [스쿼시 커밋](#스쿼시-커밋)

---

## 도움이 되는 자원

### 시작하기

- [컨트리뷰터 가이드] - 어떻게 쿠버네티스 프로젝트에 기여할 수 있는지에 대한
  가이드
- [개발자 가이드] - 쿠버네티스 프로젝트에 직접적으로 코드를 기여하기 위한 방법을 제공하는
  가이드
- [보안과 정보 공개] - 취약점 리포트와 보안 릴리스 프로세스에 대한
  가이드

### SIG 및 기타 그룹

- [마스터 그룹 목록][sigs]

### 커뮤니티

- [달력] - 쿠버네티스 커뮤니티 이벤트 모두 확인 (SIG/WG 미팅,
  이벤트 등.)
- [kubernetes-dev] - 쿠버네티스 개발 메일링 목록
- [쿠버네티스 포럼] - 쿠버네티스 공식 포럼
- [Slack 채널] - 쿠버네티스 공식 Slack
- [Stack Overflow] - 쿠버네티스 엔드-사용자의 질문에 답변할 수 있는 장소
- [YouTube 채널] - 쿠버네티스 커뮤니티 공식 채널


### 작업 흐름

- [Gubernator 대시보드] - 주목이 필요한 수신/발신 풀 리퀘스트
  확인
- [Prow] - 쿠버네티스 CI/CD 시스템
- [Tide] - 머지와 테스트를 관리하는 Prow 플러그인 [Tide 대시보드]
- [Bot 명령] - 쿠버네티스 Bot과 상호 작용하는데 사용하는 명령 (예시:
  `/cc`, `/lgtm`, `/retest`)
- [GitHub 레이블] - 쿠버네티스 프로젝트 전체에서 사용되는 레이블 목록
- [@dims]가 관리하는 [쿠버네티스 코드 검색]


### 테스트

- [Prow] - 쿠버네티스 CI/CD 시스템
- [테스트 그리드] - 테스트 기록과 관련된 정보 확인
- [Triage 대시보드] - 더 나은 문제 해결을 위해 유사한 오류를
  수집
- [Velodrome] - 작업과 테스트 상태를 추적하는 대시보드


### 중요한 이메일 별칭

- community@kubernetes.io - 커뮤니티 문제에 대한 커뮤니티 팀(SIG 컨트리뷰터 경험자)의
  누군가에게 메일을 전송
- conduct@kubernetes.io - 비공개 메일링 목록은 행동 강령 위원회에
  문의
- steering@kubernetes.io - 공개 자료와 공개 주소와 함께 운영 위원회에 메일을
  전송
- steering-private@kubernetes.io - 민감한 사항에 대해서는 운영 위원회에 개인적으로
  메일 전송
- social@cncf.io - CNCF 소셜 팀에 문의. 블로그, 트위터 계정,
  기타 소셜 요소들


### 다른 유용한 링크

- [개발자 통계] - CNCF가 관리하는 프로젝트에 대한 개발자 통계
  확인

---

## GitHub에서 효과적으로 의사 소통하기


### 서로에게 훌륭한 방법

첫 번째 단계로, [행동 강령]에 익숙해지자.


#### 좋은/나쁜 의사 소통 예시

이슈를 제기하거나, 도움을 청할 때, 요청에 예의를 갖춘다.

  🙂 “Y를 할 때 X가 컴파일되지 않습니다. 다른 의견이 있습니까?”

  😞 “X가 작동하지 않습니다! 고쳐주세요!”

PR을 종료할 때, 머지해야 할 요건을 충족하지 못하는 이유를 설명하는
구체적이고 친절한 메시지를 전달한다.

🙂 “이 기능은 유스 케이스 X를 지원할 수 없기 때문에 이 PR을 닫습니다.
    이 경우, Y 도구로 구현하는 것이 더 낫습니다.
    이 일에 도움을 주셔서 감사합니다.”

😞 “API 규약을 따르지 않는 이유는 무엇입니까? 이것은 반드시 다른 곳에서 해야합니다!”

---

## 컨트리뷰션 전송

### CLA 서명

컨트리뷰션을 전송하기 전에, 반드시 [컨트리뷰터 라이센스 동의(CLA)][cla]가 필요하다.
쿠버네티스 프로젝트는 귀하 또는 귀하의 회사가 CLA에 서명한 경우에만
_오직_ 기여할 수 있다.

CLA에 서명에 문제가 발생한다면,
[CLA 문제해결 가이드라인]을 참조하라.


### 이슈의 오픈 및 응답

GitHub 이슈는 버그 리포트, 개선 요청 또는 실패한 테스트와 같은 문제를 추적하는
주요 수단이다.
이것은 [사용자 지원 요청]을 위한 것이 **아니다**.
그 경우에는, [문제해결 가이드]를 확인하거나, [Stack Overflow]에 문제를 리포트하거나,
[쿠버네티스 포럼]에서 후속 조치를 취한다.

**참조:**
- [레이블]
- [Prow 명령][commands]


#### 이슈 생성

- 가능하다면 이슈 템플릿을 사용한다. 올바른 것을 사용하면 다른 컨트리뷰터가
  귀하의 문제를 응답하는 데 도움이 된다.
  - 이슈 템플릿 자체에 설명된 지침을 따른다.
- 발생한 이슈를 설명한다.
- 적절한 [레이블]을 지정한다. 확실하지 않은 경우, [k8s-ci-robot][prow] bot
  ([쿠버네티스 CI bot][prow])이 효과적으로 분류된 레이블로 이슈에
  답장할 것이다.
- [`/assign @<username>`][assign] 또는 [`/cc @<username>`][cc]를 사용하여 이슈를 선택적으로 할당한다.
  이슈에 많은 사람을 할당하는 것보다 올바른 레이블을 적용하는 것이
  이슈를 효과적으로 선별할 것이다.


#### 이슈 응답

- 이슈를 해결할 때, 다른 사람과의 중복 작업을 피하기 위해 작업중임을 알리는
  내용의 댓글을 작성한다.
- 나중에 언제든지 무언가를 해결할 때,
  이슈를 닫기 전에 사람들에게 알린다.
- 다른 PR 또는 이슈(또는 접근 가능한 자료)에 대한 참조를 포함한다. 예: _"ref: #1234"_.
  이는 관련 작업이 다른 곳에서 처리되었음을
  확인하기에 유용한다.


### 풀 리퀘스트 오픈

풀 리퀘스트(PR)는 git 저장소에 저장되는 코드,
문서 또는 다른 형태의 작업을 기여하는 주요 수단이다.

**참고:**
- [레이블]
- [Prow 명령][commands]
- [풀 리퀘스트 과정]
- [GitHub 작업 흐름]


#### 풀 리퀘스트 생성

- 가능한 풀 리퀘스트 템플릿의 지시를 따른다.
  템플릿은 PR에 응답하는 사람들을 도와줄 것이다.
- 깨진 링크, 오타 또는 문법 오류와 같은 [사소한 수정]이 있으면, 전체 문서에서 다른 잠재적인 실수를
  검토한다. 같은 문서에서 작은 수정을 위해
  여러개의 PR을 오픈하지 않는다.
- PR과 관련된 이슈나 PR이 해결할 수 있는 모든 이슈를 참조한다.
- 단일 커밋에서 지나치게 큰 변경은 피한다.
  대신에, PR을 여러 개의 작고 논리적인 커밋으로 나눈다.
  이것은 PR을 보다 쉽게 리뷰할 수 있도록 만든다.
- 더 자세한 설명이 필요하다고 생각하는 자신의 PR에 대한
  의견을 말한다.
- [`/assign @<username>`][assign]으로 PR을 선택적으로 할당한다.
  많은 리뷰어를 지정하더라도, 더 빨리 PR을 검토되지는 않는다.
- PR이 _"진행중인 작업"_ 으로 간주되면, 접두사로 `[WIP]`를 붙이거나,
  [`/hold`][hold] 명령을 사용한다. 이는 `[WIP]` 또는 hold가 해제될 때까지
  PR이 합쳐지는 것을 막을 것이다.
- 만약 PR이 리뷰되지 않았다면, 동일한 변경 사항으로 새로운 PR을 열고 닫지 않는다.
  `@<github username>` 명령을 사용하여 리뷰어에게 Ping한다.


#### PR 설명 예시

```
Ref. #3064 #3097
All files owned by SIG testing were moved from `/devel` to the new folder `/devel/sig-testing`.
(SIG 테스트에서 소유한 모든 파일을 `/devel`에서 새폴더 `/devel/sig-testing`으로 이동하였습니다.)

/sig contributor-experience
/cc @stakeholder1 @stakeholder2
/kind cleanup
/area developer-guide
/assign @approver1 @approver2 @approver3
```

PR에 무엇이 포함되어 있는가?
- **라인 1** - 다른 이슈나 PR의 레퍼런스 (#3064 #3097).
- **라인 2** - PR에서 무엇이 수행되었는지 간단한 설명한다.
- **라인 4** - `/sig contributor-experience` [명령][commands]으로
  [SIG][sigs] 할당한다.
- **라인 5** - 특정 이슈 또는 PR에 관심을 가지고 있는 특정한 리뷰어는 [`/cc`][cc] 명령으로
  지정한다.
- **라인 6** - [`/kind cleanup`][kind] 명령은 코드, 프로세스 또는 기술적 부재 정리와 관련하여
  PR을 분류하는 [레이블][레이블]을
  추가한다.
- **라인 7** - [`/area developer-guide`][kind] 명령은 개발자 가이드와 관련된 이슈 또는
  PR을 분류한다.
- **라인 8** - [`/assign`][assign] 명령은 승인자를 PR에 할당한다.
  승인자는 [k8s-ci-robot][prow]에 의해 제안되며,
  [OWNERS] 파일의 소유자 목록에서 선택된다.
  그들은 PR을 검토한 후에, [`/approve`][approve] 레이블을 추가할 것이다.


#### 풀 리퀘스트 문제 해결

PR이 제안된 후, 쿠버네티스 CI 플랫폼 [Prow]에 의해 일련의 테스트가 실행된다.
테스트가 실패하면, [k8s-ci-robot][prow]은 실패한 테스트 및 사용 가능한 로그에 대한
링크를 PR에서 응답한다.

새로운 커밋을 PR에 푸시하면 자동으로 테스트가 다시 실행된다.

가끔 쿠버네티스 CI 플랫폼에 문제가 있을 수도 있다.
컨트리뷰션이 모든 로컬 테스트를 통과하더라도 다양한 이유로 이슈가 발생할 수 있다.
`/retest` 명령으로 테스트를 재실행 할 수 있다.

특정 테스트 문제 해결에 대한 자세한 내용은, [테스트 가이드]를 참고하라.


### 레이블

쿠버네티스는 이슈와 풀 리퀘스트를 분류하고 선별하기 위해 [레이블]을 사용한다.
올바른 레이블을 적용하면 이슈 또는 PR이 더 많이 효과적으로 검토하는데
도움을 줄 수 있다.

**참고:**
- [레이블]
- [Prow 명령][commands]

자주 사용되는 레이블:
- [`/sig <sig name>`][kind] 이슈나 PR의 소유권을 [SIG][SIGs]로
  지정.
- [`/area <area name>`][kind] 이슈 또는 PR을 특정 [area][레이블]과
  연결.
- [`/kind <category>`][kind] 이슈나 PR을 [분류][레이블].

---

## 로컬에서 작업

풀 리퀘스트 작업을 제안하기 전에, 로컬에서 일정 수준의 작업을 수행해야 한다.
git에 익숙하지 않다면, [Atlassian git 튜토리얼]이 좋은 출발점이 될 것이다.
대안으로는, Stanford의 [Git magic] 튜토리얼이 좋은
다국어 옵션이다.

**참고:**
- [Atlassian git 튜토리얼]
- [Git magic]
- [GitHub 작업 흐름]
- [로컬 테스트]
- [개발자 가이드]


### 브랜치 전략

쿠버네티스는 GitHub의 표준인 _"포크와 풀"_ 작업 흐름을 사용한다.
git 용어로는, 개인적인 포크는 _"`origin`"_ 으로, 실제 프로젝트의 저장소는
_"`upstream`"_ 으로 부른다.
개인 브랜치(`origin`)를 프로젝트(`upstream`)의 최신 상태로 유지하려면,
로컬 작업 복사본 내에 반드시 구성되어야 한다.


#### Upstream 추가

원격지로 `upstream`을 추가하고, 푸시할 수 없도록 설정한다.

```
# replace <upstream git repo> with the upstream repo url
# example:
#  https://github.com/kubernetes/kubernetes.git
#  git@github.com/kubernetes/kubernetes.git

git remote add upstream <upstream git repo>
git remote set-url --push upstream no_push
```

`git remote -v` 명령을 수행함으로써 원격지로 설정되어진 목록을
나열할 수 있다.


#### Fork를 동기화 상태로 유지하기

`master` 브랜치에서 `upstream`과 _"rebase"_ 로부터 모든 변경 사항을 가져온다.
이는 로컬 저장소와 `upstream` 프로젝트를 동기화할 것이다.

```
git fetch upstream
git checkout master
git rebase upstream/master
```

기능이나 수정 작업을 위해 새로운 브랜치를 만들기 이전에
작업을 최소한으로 수행해야만 한다.

```
git checkout -b myfeature
```

#### 스쿼시 커밋

[스쿼시 커밋]의 주요 목적은 변경된 내용을 명백하게 읽을 수 있는 git 히스토리 또는 로그를 생성하는 것이다.
보통 이것은 PR 정정의 마지막 단계에서 수행된다.
커밋을 스쿼시해야 하는지 확실하지 않다면,
더 많은 것을 가지고 있는 측에서 잘못을 범하는 것이 더 좋기 때문에,
PR을 검토하고 승인하도록 지정된 다른 참여자의 판단에 맡긴다.



[컨트리뷰터 가이드]: /contributors/guide/README.md
[개발자 가이드]: /contributors/devel/README.md
[Gubernator 대시보드]: https://gubernator.k8s.io/pr
[prow]: https://prow.k8s.io
[tide]: http://git.k8s.io/test-infra/prow/cmd/tide/pr-authors.md
[tide 대시보드]: https://prow.k8s.io/tide
[bot 명령]: https://go.k8s.io/bot-commands
[gitHub 레이블]: https://go.k8s.io/github-labels
[쿠버네티스 코드 검색]: https://cs.k8s.io/
[@dims]: https://github.com/dims
[달력]: https://calendar.google.com/calendar/embed?src=cgnt364vd8s86hr2phapfjc6uk%40group.calendar.google.com
[kubernetes-dev]: https://groups.google.com/forum/#!forum/kubernetes-dev
[Slack 채널]: http://slack.k8s.io/
[Stack Overflow]: https://stackoverflow.com/questions/tagged/kubernetes
[YouTube 채널]: https://www.youtube.com/c/KubernetesCommunity/
[triage 대시보드]: https://go.k8s.io/triage
[테스트 그리드]: https://testgrid.k8s.io
[velodrome]: https://go.k8s.io/test-health
[개발자 통계]: https://k8s.devstats.cncf.io
[행동 강령]: /code-of-conduct.md
[사용자 지원 요청]: /contributors/guide/issue-triage.md#determine-if-its-a-support-request
[문제해결 가이드]: https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/
[쿠버네티스 포럼]: https://discuss.kubernetes.io/
[풀 리퀘스트 과정]: /contributors/guide/pull-requests.md
[github 작업 흐름]: /contributors/guide/github-workflow.md
[prow]: https://git.k8s.io/test-infra/prow#prow
[cla]: /CLA.md#how-do-i-sign
[CLA 문제해결 가이드라인]: /CLA.md#troubleshooting
[commands]: https://prow.k8s.io/command-help
[kind]: https://prow.k8s.io/command-help#kind
[cc]: https://prow.k8s.io/command-help#cc
[hold]: https://prow.k8s.io/command-help#hold
[assign]: https://prow.k8s.io/command-help#assign
[SIGs]: /sig-list.md
[테스트 가이드]: /contributors/devel/sig-testing/testing.md
[레이블]: https://git.k8s.io/test-infra/label_sync/labels.md
[사소한 수정]: /contributors/guide/pull-requests.md#10-trivial-edits
[GitHub 작업 흐름]: /contributors/guide/github-workflow.md#3-branch
[스쿼시 커밋]: /contributors/guide/pull-requests.md#6-squashing-and-commit-titles
[owners]: /contributors/guide/owners.md
[로컬 테스트]: /contributors/guide/README.md#testing
[Atlassian git 튜토리얼]: https://www.atlassian.com/git/tutorials
[git magic]: http://www-cs-students.stanford.edu/~blynn/gitmagic/
[보안과 정보 공개]: https://kubernetes.io/docs/reference/issues-security/security/
[approve]: https://prow.k8s.io/command-help#approve
