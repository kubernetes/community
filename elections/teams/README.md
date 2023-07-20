# Elections for Kubernetes Teams

As Kubernetes now has access to self-hosted preference election software,
we can host any number of elections for community groups, as they need them.

## Who Can Request an Election

Any "Kubernetes team" can request an election.  This includes SIGs, Working 
Groups, operational teams like Release Engineering or Enhancements, 
Subprojects, and even wholly subsidiary projects like Prow and ClusterAPI.
Basically, if your group is all Kubernetes contributors, you can ask for an
election.

This does not include CNCF projects that are not components of Kubernetes and
subject to Kubernetes governance.  At some point, the CNCF may offer 
elections-as-a-service for those other projects.

## How It Works

Kubernetes uses an election tool called [Elekto], an instance of which is 
hosted at [elections.k8s.io].  Your team's election gets added as metadata
to the kubernetes/community repository, and that automatically causes 
an election to appear in the app. 

You assign some of your team to be election administrators, and they are
in charge of running the election.  You also assemble and supply a list
of GithubIDs for valid voters in the election, and the dates that candidate
statements are due, followed by when voting opens and closes.  Candidates
join the election by publishing candidate profiles to your election directory.

Depending on availability and the nature of your election, Contributor Comms 
may help publicize your election and remind voters of deadlines. Otherwise,
you will do this.

Once voting has closed, one of your election administrators can calculate
the result of the election, and publish it either through the app or directly
to your team's channel or mailing list.

## Requesting by Issue

If you are not already familiar with [Elekto], you should request an election
by filing an [election request issue].  Please fill out all of the information
in the template; it's required for the election.

A member of the Elections Subproject will contact you and help you through
setting up the election metadata.  Please allow at least a week before your election
needs to start.

## Requesting by Pull Request

If you are familiar with [Elekto] from prior elections, then you may create
your own metadata files and submit your election as a pull request, which 
will make it happen faster.

Each team election goes in its own folder in the kubernetes/community repo,
using the path `elections/teams/team-name/election-name`, 
e.g. `elections/teams/clusterapi/leads-2022`.  This folder should then have
the following files in it, with complete contents.  See the 
[Elekto administration docs] for more details.

* election.yaml file defining the election
* election-desc.md file with a text description of the election
* voters.yaml with the initial list of voters

The easiest path is to copy these files from another election and then 
modify them. The above will require you to have already decided your list of 
Election Admins and the dates for your election.

A member of the Elections Subproject will then review your PR, offer any required
data corrections, and approve it.

## Additional Notes

* Candidates do not have to be people; sometimes you may want to run an
  "election" between development alternatives or graphic design options.
  However, consider doing a survey for those instead.
* Voters and Admins do not have to be Kubernetes Org Members.  
  They do have to have GitHub IDs.


[Elekto]: https://elekto.dev
[elections.k8s.io]: https://elections.k8s.io
[preference elections]: https://en.wikipedia.org/wiki/Preferential_voting
[election request issue]: /issues/new/choose
[Elekto administration docs]: https://elekto.dev/docs/administration/
