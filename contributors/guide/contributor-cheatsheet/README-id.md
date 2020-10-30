<!-- omit in toc -->
# _Cheat Sheet_ Kontributor Kubernetes

Kumpulan _resources_ umum yang digunakan ketika berkontribusi ke Kubernetes, 
termasuk _tips_, trik, dan _best practices_ yang digunakan di dalam 
proyek Kubernetes. Ini merupakan referensi singkat ("TL;DR") informasi 
yang bermanfaat untuk meningkatkan pengalaman kamu ketika berkontribusi 
di GitHub menjadi lebih baik.

**Daftar Isi**
- [Sumber Penting](#sumber-penting)
  - [Mulai Berkontribusi](#mulai-berkontribusi)
  - [_SIG_ dan Grup Lainnya](#sig-dan-grup-lainnya)
  - [Komunitas](#komunitas)
  - [_Workflow_](#workflow)
  - [_Testing_](#testing)
  - [Alamat Email Penting](#alamat-email-penting)
  - [Tautan Lain](#tautan-lain)
- [Berkomunikasi Secara Efektif di GitHub](#berkomunikasi-secara-efektif-di-github)
  - [Bagaimana Cara Bekerja Sama dengan Baik](#bagaimana-cara-bekerja-sama-dengan-baik)
    - [Contoh Komunikasi Yang Baik/Buruk](#contoh-komunikasi-yang-baikburuk)
- [Mengumpulkan Kontribusi](#mengumpulkan-kontribusi)
  - [Menandatangani CLA](#menandatangani-cla)
  - [Membuka dan Menanggapi Isu](#membuka-dan-menanggapi-isu)
    - [Membuat Sebuah Isu](#membuat-sebuah-isu)
    - [Menanggapi sebuah Isu](#menanggapi-sebuah-isu)
  - [Membuka sebuah Pull Request (PR)](#membuka-sebuah-pull-request-pr)
    - [Membuat sebuah Pull Request (PR)](#membuat-sebuah-pull-request-pr)
    - [Contoh Deskripsi PR](#contoh-deskripsi-pr)
    - [_Troubleshooting_ sebuah PR](#troubleshooting-sebuah-pr)
  - [Label](#label)
- [Bekerja pada Mesin Lokal](#bekerja-pada-mesin-lokal)
  - [Mekanisme _Branch_](#mekanisme-branch)
    - [Menambahkan _Upstream_](#menambahkan-upstream)
    - [Menjaga agar _Fork_ Kamu tetap Sinkron](#menjaga-agar-fork-kamu-tetap-sinkron)
    - [Melakukan _Commit Squashing_](#melakukan-commit-squashing)

---

## Sumber Penting

### Mulai Berkontribusi

- [Panduan Kontributor] - Panduan bagaimana cara berkontribusi dalam proyek Kubernetes.
- [Panduan Pengembang] - Panduan untuk berkontribusi dalam pengembangan kode pada proyek Kubernetes.
- [Informasi _Security_ dan _Disclosure_] - Panduan pelaporan celah keamanan 
  dan proses rilis keamanan.

### _SIG_ dan Grup Lainnya

- [_List_ Grup Master][sigs]

### Komunitas

- [Kalender] - Melihat semua acara Komunitas Kubernetes View (pertemuan SIG/WG,
  acara, dll.)
- [kubernetes-dev] - Alamat email pengembangan Kubernetes
- [Forum Kubernetes] - Forum resmi Kubernetes.
- [Slack _channels_] - Slack resmi Kubernetes.
- [Stack Overflow] - _Platform_ tanya jawab pengguna Kubernetes.
- [YouTube _Channel_] - _Channel_ resmi untuk komunitas Kubernetes.


### _Workflow_

- [Dasbor Gubernator] - Melihat Pull Request yang masuk dan keluar yang memerlukan perhatian.
- [Prow] - Mekanisme CI/CD Kubernetes.
- [Tide] - _Plugin_ Prow yang melakukan manajemen _merge_ dan _test_. [Dasbor Tide]
- [Perintah Bot] - Perintah yang dapat kamu gunakan untuk berinteraksi dengan Bot Kubernetes (contoh:
  `/cc`, `/lgtm`, dan `/retest`)
- [Label GitHub] - _List label_ yang digunakan pada proyek Kubernetes
- [Pencarian Kode Kubernetes], di-_maintain_ oleh [@dims]

### _Testing_

- [Prow] - Mekanisme CI/CD Kubernetes.
- [Test Grid] - Melihat data _historical testing_ beserta informasi terkait.
- [Dasbor Triase] - Melakukan agregasi _failure_ untuk mekanisme _troubleshoot_ yang lebih baik.


### Alamat Email Penting

- community@kubernetes.io - Mengirim surat pada seseorang yang berada dalam tim komunitas (Kontributor SIG
  _Experience_) mengenai isu komunitas.
- conduct@kubernetes.io - Mengirim surat pada komite _Code of Conduct_, _mailing list private_.
- github@kubernetes.io - Mengirim Surat pada Tim Administrasi GitHub secara pribadi, untuk item sensitif.
- steering@kubernetes.io - Mengirim surat pada _steering committee_. Alamat publik dengan _archieve_.
- steering-private@kubernetes.io - Mengirim surat pada _steering committee_ secara _private_, untuk 
  informasi yang sensitif.
- social@cncf.io - Mengirim surat pada tim CNCF; blog, akun twitter, and
  properti sosial lainnya.


### Tautan Lain

- [Statistik Pengembang] - Melihat statistik pengembang untuk semua proyek yang dikelola oleh CNCF.
- [Rilis Patch Kubernetes] Jadwal dan informasi kontak tim untuk rilis _patch_ Kubernetes

---

## Berkomunikasi Secara Efektif di GitHub


### Bagaimana Cara Bekerja Sama dengan Baik

Pada tahap awal, pelajari dan pahami [Code of Conduct].


#### Contoh Komunikasi Yang Baik/Buruk

Ketika membuka sebuah isu, atau mencari bantuan, tolong bersikap dengan sopan ketika melakukan hal tersebut:

  🙂 “X tidak dapat dikompilasi, apakah kamu memiliki saran?”

  😞 “X tidak bekerja sebagaimana mestinya! Tolong perbaiki!”

Ketika menutup sebuah PR, berikan penjelasan yang rinci mengenai alasan kenapa PR 
tersebut tidak memenuhi standar untuk di-_merge_.

🙂 “Aku menutup PR ini karena fitur ini tidak mendukung kebutuhan X. 
    Sesuai kesepakatan, akan lebih baik jika ini diimplementasikan dengan Y. 
    Terima kasih sudah menyempatkan waktu untuk mengerjakan hal ini."

😞 “Mengapa PR ini tidak mengikuti konvensi API? 
    Hal ini harusnya tidak dilakukan di sini!"

---

## Mengumpulkan Kontribusi

### Menandatangani CLA

Sebelum kamu mengumpulkan kontribusi kamu, kamu harus terlebih dahulu [menyetujui _Contributor License Agreement(CLA)_][cla]. Proyek Kubernetes _hanya_ menerima kontribusi yang kamu kerjakan apabila kamu sudah menyetujui CLA.

Apabila kamu kesulitan ketika menyetujui CLA, ikuti [petunjuk _troubleshooting_ CLA].


### Membuka dan Menanggapi Isu

Isu GitHub merupakan mekanisme _tracking_ berbagai hal yang ada, termasuk pelaporan _bug_, 
permintaan peningkatan fitur, atau pelaporan isu lainnya seperti terjadi kegagalan ketika menjalankan 
_test_. Hal tersebut **tidak** diperuntukkan bagi [_user support request_]. Untuk tujuan tersebut, 
kamu bisa membaca [petunjuk _troubleshooting_], laporkan permasalahan yang ada ke [Stack Overflow] 
atau ikuti [forum Kubernetes].

**Referensi**
- [Label]
- [Perintah Prow][perintah bot]


#### Membuat Sebuah Isu

- Gunakan templat isu (jika tersedia). Menggunakan templat yang tersedia akan 
  memudahkan kontributor lain ketika menanggapi isu yang kamu buat.
  - Ikuti petunjuk yang dideskripsikan di templat tersebut.
- Berikan deskripsi yang cukup ketika membuat suatu isu.
- Gunakan [label] yang tepat. Jika kamu kurang yakin, [k8s-ci-robot][prow] bot
  ([Kubernetes CI bot][prow]) akan membalas isu yang kamu buat dengan respons
  `needed labels`.
- Selektiflah ketika meng-_assign_ suatu isu menggunakan [`/assign @<username>`][assign] atau 
  [`/cc @<username>`][cc]. Isu yang kamu buat akan ditriase secara lebih efektif 
  apabila kamu menambahkan label yang tepat.


#### Menanggapi sebuah Isu

- Ketika menghadapi sebuah isu, berikan komentar terhadap isu tersebut yang menandakan kamu 
  sedang mengerjakan isu tersebut untuk mencegah pengerjaan berulang.
- Ketika kamu sudah menyelesaikan hal tersebut, berikan komentar yang mengindikasikan 
  kamu sudah menyelesaikannya sebelum kamu menutup isu tersebut, hal ini bertujuan agar 
  orang lain tahu alasan kenapa kamu menutup isu tersebut.
- Masukkan referensi ke PR atau isu lain (atau material lain yang dapat diakses)  
  misalnya _"ref: #1234"_. Hal ini nantinya dapat digunakan untuk mengidentifikasi 
  hal terkait yang mungkin sudah dibahas/dikerjakan di tempat lain.


### Membuka sebuah Pull Request (PR)

Pull requests (PR) adalah mekanisme utama yang digunakan untuk melakukan kontribusi kode, 
dokumentasi, atau segala bentuk hal yang disimpan dalam repositori git.

**Referensi:**
- [Label]
- [Perintah Prow][perintah bot]
- [Pull request process]
- [GitHub workflow]


#### Membuat sebuah Pull Request (PR)

- Ikuti petunjuk yang ada pada templat PR (jika tersedia). Menggunakan templat yang tersedia akan 
  memudahkan kontributor lain ketika menanggapi PR yang kamu buat.
- Jika hal tersebut merupakan [_trivial fix_] seperti tautan yang _broken_, 
  _typo_ atau kesalahan _grammar_, _review_ dokumen secara menyeluruh untuk mengecek 
  apakah terjadi kesalahan yang serupa. Jangan membuka multipel PR untuk 
  _fix_ kecil pada dokumen yang sama.
- Berikan referensi ke isu yang terkait dengan PR kamu, atau isu lain yang mungkin 
  dapat diselesaikan dengan PR yang kamu buat.
- Hindari perubahan yang sangat besar dalam suatu _commit_, Sebaliknya, 
  pisahkan PR kamu dalam multipel PR yang lebih kecil. Ini akan memudahkan proses _review_ 
  PR yang kamu buat.
- Berikan komentar pada PR yang kamu buat, apabila kamu merasa ada hal yang membutuhkan 
  penjelasan lebih lanjut.
- Selektiflah ketika meng-_assign_ suatu isu menggunakan [`/assign @<username>`][assign]. 
  Meng-_assign reviewer_ secara berlebihan tidak mempercepat proses _review_ yang dilakukan 
  untuk PR kamu.
- Jika PR kamu masih dalam tahap _"Work in progress"_ berikan prefiks `[WIP]`
  atau gunakan perintah [`/hold`][hold]. Hal ini mencegah agar PR tidak di-_merge_ hingga _WIP_ 
  dihilangkan.
- Jika PR kamu tidak di-_review_, jangan tutup PR tersebut dan membukan PR lain 
  dengan perubahan yang sama. _Ping reviewer_ PR kamu dengan komentar `@<github username>`.


#### Contoh Deskripsi PR

```
Ref. #3064 #3097
Semua file yang dimiliki oleh SIG testing dipidahkan dari folder `/devel` ke folder `/devel/sig-testing`.

/sig contributor-experience
/cc @stakeholder1 @stakeholder2
/kind cleanup
/area developer-guide
/assign @approver1 @approver2 @approver3
```

Apa saja yang dimasukkan dalam deskripsi:
- **Baris 1** - Referensi ke isu atau PR lain (#3064 #3097).
- **Baris 2** - Deskripsi singkat mengenai apa yang dikerjakan dalam PR tersebut.
- **Baris 4** - [SIG][sigs] _assignment_ dengan [perintah]
  `/sig contributor-experience`..
- **Baris 5** - _Reviewer_ yang terkait dengan isu atau PR yang dispesifikasikan dengan perintah [`/cc`][cc].
- **Baris 6** - Perintah [`/kind cleanup`][kind] menambahkan [label] yang 
  mengkategorisasi isu atau PR yang terkait untuk membersihkan kode, _process_, atau _technical
  debt_.
- **Baris 7** - Perintah [`/area developer-guide`][kind] mengkategorisasi isu atau PR 
  yang terkait dengan petunjuk pengembang.
- **Baris 8** - Perintah [`/assign`][assign] meng-_assign_ seorang _approver_ untuk PR.
  Seorang _approver_ akan disarankan oleh [k8s-ci-robot][prow] dan dipilih dari _list_ 
  _owner_ pada _file_ [OWNERS]. Merekalah yang akan menambahkan _label_ 
  [`/approve`][approve] pada PR yang sudah di-_review_.


#### _Troubleshooting_ sebuah PR

Setelah PR kamu diajukan, serangkaian _testing_ akan dijalankan oleh platform CI Kubernetes, [Prow]. 
Jika terdapat salah satu _test_ yang gagal, maka [k8s-ci-robot][prow] akan memberikan 
balasan pada PR kamu beserta tautan yang memberikan _log_ dari _testing_ yang gagal dijalankan.

Apabila kamu mem-_push commit_ baru, _test_ pada PR kamu akan secara otomatis di-_trigger_.

Terkadang, bisa jadi terdapat masalah pada platform CI Kubernetes. 
Hal ini dapat terjadi karena berbagai alasan bahkan ketika _test_ yang kamu jalankan di 
mesin lokal kamu berhasil. Kamu dapat men-_trigger_ ulang _test_ dengan cara memanggil perintah 
`/retest`.

Untuk informasi lebih lanjut, baca [Panduan _Testing_].


### Label

Kubernetes menggunakan [label] untuk melakukan kategorisasi 
dan triase isu dan PR. Penggunaan label yang benar akan membuat triase 
pada isu atau PR yang kamu ajukan menjadi lebih efektif.

**Referensi:**
- [Label]
- [Perintah Prow]

Label yang sering digunakan:
- [`/sig <sig name>`][kind] Meng-_assign_ [SIG][SIGs] yang bertindak sebagai _owner_ 
  sebuah isu atau PR.
- [`/area <area name>`][kind] Mengasosiasikan [area][label].
- [`/kind <category>`][kind] [Mengkategorisasikan][label] isu atau PR.

---

## Bekerja pada Mesin Lokal

Sebelum kamu membuat sebuah PR, sebagian besar kamu akan mengerjakan pekerjaan kamu pada mesin lokal. 
Jika kamu merupakan pengguna baru git, [Tutorial git Atlassian ] merupakan awal 
pembelajaran yang baik. Sebagai alternatif lain, juga terdapat tutorial Stanford's [Git magic] untuk pilihan multi-bahasa yang bagus.

**Referensi:**
- [Tutorial git Atlassian]
- [Git magic]
- [GitHub workflow]
- [Testing locally]
- [Panduan Pengembang]


### Mekanisme _Branch_

Proyek Kubernetes menggunakan mekanisme _"Fork and Pull"_ yang merupakan 
standar GitHub. Dalam terminologi git, _fork_ yang kamu buat disebut sebagai _"`origin`"_ 
dan git proyek yang sebenarnya disebut sebagai _"`upstream`"_. Untuk menjaga _branch_ 
(`origin`) tetap _up to date_ dengan proyek (`upstream`), _branch_ tersebut harus dikonfigurasikan 
pada _copy_ lokal.


#### Menambahkan _Upstream_

Tambahkan `upstream` sebagai _remote_, dan atur agar kamu tidak dapat mem-_push_ ke sana.

```
# ganti <upstream repositori git> dengan url upstream repositori
# contoh:
#  https://github.com/kubernetes/kubernetes.git
#  git@github.com/kubernetes/kubernetes.git

git remote add upstream <upstream git repo>
git remote set-url --push upstream no_push
```

Versifikasi langkah ini dapat dilakukan dengan cara menjalankan 
`git remote -v` yang selanjutnya akan menampilah seluruh 
_remote_ yang sudah kamu atur.


#### Menjaga agar _Fork_ Kamu tetap Sinkron

_Fetch_ semua perubahan dari `upstream` dan lakukan _"rebase"_ pada `master` _branch_ lokal kamu. 
Dengan demikian repositori lokal kamu akan tertap sinkron dengan proyek `upstream`.

```
git fetch upstream
git checkout master
git rebase upstream/master
```

Kamu setidaknya harus melakukan hal ini sebelum membuat sebuah _branch_ baru 
yang akan kamu gunakan untuk mengerjakan fitur baru atau melakukan _fix_.

```
git checkout -b myfeature
```

#### Melakukan _Commit Squashing_

Tujuan utama dari [_commit squashing_] adalah untuk membuat 
histori atau _log_ git yang mudah dibaca dan bersih. Biasanya hal ini 
dilakukan pada fase akhir dari revisi yang kamu buat. Jika kamu masih belum yakin apakah kamu 
harus melakukan _commit squashing_ atau tidak, biarkan revisi kamu apa adanya sampai 
ada saran khusus dari kontributor yang di-_assign_ untuk me-_review_ atau meng-_approve_ revisi kamu apakah _commit_ 
_squashing_ perlu dilakukan atau tidak.

[Panduan Kontributor]: /contributors/guide/README.md
[Panduan Pengembang]: /contributors/devel/README.md
[dasbor gubernator]: https://gubernator.k8s.io/pr
[prow]: https://prow.k8s.io
[tide]: http://git.k8s.io/test-infra/prow/cmd/tide/pr-authors.md
[dasbor tide]: https://prow.k8s.io/tide
[perintah bot]: https://go.k8s.io/bot-commands
[Label GitHub]: https://go.k8s.io/github-labels
[Pencarian Kode Kubernetes]: https://cs.k8s.io/
[@dims]: https://github.com/dims
[kalender]: https://calendar.google.com/calendar/embed?src=calendar%40kubernetes.io
[kubernetes-dev]: https://groups.google.com/forum/#!forum/kubernetes-dev
[slack _channels_]: http://slack.k8s.io/
[Stack Overflow]: https://stackoverflow.com/questions/tagged/kubernetes
[youtube _channel_]: https://www.youtube.com/c/KubernetesCommunity/
[dasbor triase]: https://go.k8s.io/triage
[test grid]: https://testgrid.k8s.io
[Statistik Pengembang]: https://k8s.devstats.cncf.io
[code of conduct]: /code-of-conduct.md
[_user support request_]: /contributors/guide/issue-triage.md#determine-if-its-a-support-request
[petunjuk _troubleshooting_]: https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/
[forum kubernetes]: https://discuss.kubernetes.io/
[pull request process]: /contributors/guide/pull-requests.md
[github workflow]: /contributors/guide/github-workflow.md
[prow]: https://git.k8s.io/test-infra/prow#prow
[cla]: /CLA.md#how-do-i-sign
[petunjuk _troubleshooting_ cla]: /CLA.md#troubleshooting
[perintah]: https://prow.k8s.io/command-help
[kind]: https://prow.k8s.io/command-help#kind
[cc]: https://prow.k8s.io/command-help#cc
[hold]: https://prow.k8s.io/command-help#hold
[assign]: https://prow.k8s.io/command-help#assign
[SIGs]: /sig-list.md
[testing guide]: /contributors/devel/sig-testing/testing.md
[label]: https://git.k8s.io/test-infra/label_sync/labels.md
[_trivial fix_]: /contributors/guide/pull-requests.md#10-trivial-edits
[GitHub workflow]: /contributors/guide/github-workflow.md#3-branch
[_commit squashing_]: /contributors/guide/pull-requests.md#6-squashing-and-commit-titles
[_owner_]: /contributors/guide/owners.md
[testing locally]: /contributors/guide/README.md#testing
[Tutorial git Atlassian]: https://www.atlassian.com/git/tutorials
[git magic]: http://www-cs-students.stanford.edu/~blynn/gitmagic/
[Panduan _Testing_]: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-testing/testing.md
[Informasi _Security_ dan _Disclosure_]: https://kubernetes.io/docs/reference/issues-security/security/
[approve]: https://prow.k8s.io/command-help#approve
[Rilis Patch Kubernetes]: https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md
