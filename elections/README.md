# Kubernetes Elections

Welcome to the Kubernetes Elections area.  This directory defines the Elections
Subproject, including its products of the annual [Steering Committee election]
and elections-as-a-service for our community.  Elections is a subproject of
the [Contributor Experience] SIG.

## Steering Elections

The Elections subproject is responsible for the annual [Steering Committee election],
which is directly administered by the Election Officers.

* If you are looking to vote in the steering election, go to the [elections app].
* If you are looking for detailed information on the steering elections,
  visit [Steering Committee election] page.

## How to Request an Election for your SIG/WG

All teams in Kubernetes may use our [elections app] to run preference elections
for their team.  See the instructions on [requesting an election].

## The Elections Subproject

The Elections Subproject is a team within Kubernetes [SIG Contributor Experience]
with the following responsibilities:

* Maintain and update the elections documentation and messaging templates
* Assist K8s-Infra team in maintaining the elections software and service
* Assisting and approving SIGs/WGs in running minor elections
* Recommending a slate of Election Officers for each Steering Election

### Members

Anyone can contribute to the elections subproject. In our [OWNERS file] you
can find the list of our current approvers and reviewers.

Approvers and reviewers are also responsible for the Elections roadmap, 
maintenance, and security.  As such, new approvers must be approved by
the [Contributor Experience] chairs or by the Steering Committee.

### Communications

The elections subproject can be reached by the following mechanisms:

* tagging `/area elections` in issues or PRs in the Kubernetes/Community repo
* #sig-contribex channel in Kubernetes slack
* The regular [Contributor Experience] meeting

### Documentation

The subproject is responsible for making sure that all elections documentation
is completed and kept up-to-date.  This includes:

* Community documentation on the SC election
* Election Officer documentation on how to run an SC election
* Election messaging templates
* Elekto documentation

This documentation may include setting policy for some aspects of the Steering
election.  In those cases, the subproject is responsible for getting the SC's
approval on such items.

### Software

Elections in Kubernetes run on [Elekto].  The subproject is responsible for
maintaining these by working together with [K8s-Infra] team.  This includes
upgrades, migrations, assisting community members using the software, and handling
security reports.  Should there be a reason to change software, the subproject
will prepare recommendations for the Steering Committee to approve. It is also
responsible for any necessary scripts, such as how to pull a voter list.

### Recommending Election Officers

The Elections Subproject will be responsible for finding and recommending 
Election Officers to run the next Steering Committee election. The schedule 
for that should be:

* Early June: contact last year's EOs and determine who will be returning.
* Mid-June: put out a call within SIG-Contribex to find out who is interested
  in being an EO. Contact likely individuals 1-on-1
* Early July: submit a recommended slate of EOs to the Steering Committee,
  including alternates if possible.
* Mid-July: SC approves the EOs.
* Late July: EOs set the schedule for the election.

Election Officers should be chosen from among regular, trusted contributors to the 
Kubernetes project, with an eye towards employer, demographic, and geographic
diversity.  Further requirements can be found in the [Steering Election documentation].

Election Officers are considered part of the Elections Subproject.

### Minor Elections

The subproject will assist Kubernetes teams/SIGs/WGs in preparing any internal
elections they want to run.  This includes watching for election issues and PRs
and either helping create the appropriate files or auditing them for the teams.
The subproject will work with the Contributor Comms team to promote the elections
as appropriate.

Any named Kubernetes team (see [requesting an election]) may run an election 
in elections.k8s.io.  Since Elekto allows running multiple elections 
concurrently, the main limitation for running additional elections is the time 
of Contribex volunteers assist with the process.

### Other Responsibilities

The subproject team will make a biweekly report to SIG-Contribex's regular
meeting.

[Contributor Experience]: /sig-contributor-experience/README.md
[Elekto]: https://elekto.dev
[OWNERS file]: OWNERS
[Steering Committee election]: https://git.k8s.io/steering/elections.md
[elections app]: https://elections.k8s.io
[requesting an election]: teams/README.md
[K8s-Infra]: /sig-k8s-infra/README.md
[Steering Election documentation]: /elections/steering/README.md
