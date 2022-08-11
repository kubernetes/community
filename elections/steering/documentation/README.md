# Steering Election HOWTO

This documentation explains how to run a steering election.  As a work in
progress, it replaces older documentation where present; for topics not
covered consult the older documents in each election folder.

## Documentation contents

* This guide, which covers all process and roles.
* A folder of templates, primarily for election-related communications
* Current SQL query or devstats report for selecting voters

## Roles

### The Elections Subproject

Members of the [Elections Subproject](elections/README.md) have three 
responsibilities around steering elections:

1. Recommending Election Officers to the Steering Committee
2. Intervening if anything unexpected happens with the EOs
3. Following up on recommendations from the election retro

### The Election Officers

The Election Officers are three trusted contributors who will run one or two
Steering elections.  They are responsible for making sure that the election
happens and completes satisfactorily.

Election Officers must meet the following requirements:

* Org member for more than one year
* Eligible to vote in the election
* Pledge to administer the election without partiality to any employer, SIG, 
  or personal preference
* Tentatively available for two elections in a row (see below)
* Tentatively available for special elections in the upcoming year (see below)

Additionally, the Elections Subproject will choose Election Officers partly to
respresent the diversity of Kubernetes contributors, selecting different
genders, geographic regions, and ethnicities where possible, in order to 
avoid the appearance of bias in the elections process.  Particularly, the
three Officers must each work for a different employer.

Election Officers are responsible for:

* Planning the election, including creating a draft timeline
* Generating the voter list
* Setting up the election in the voting system
* Deciding on exception requests
* Determining candidate eligibility
* Assisting candidates with bios
* Publicizing the election in order to maximize participation
* Finalizing and reporting the election results
* Hosting an election retro & documenting changes
* Contributing to the SC election documentation

These responsibilities are detailed in the election procedures.  They collectively
fall in to four main areas, timewise:

1. Communications with candidates and voters (together with the Comms Liaison)
2. Administering the election software (together with the Infra Liaison)
3. Managing nominations and candidates
4. Responding to exception requests

Usually, the three Election Officers divide up the above major responsibilities
among the team, each Officer taking one or two of them.

Each year, at least one Election Officer must be a prior Officer in order to
ensure continuity of knowledge. As such, Officers should be theoretically 
available to do the elections two years in a row.  Further, should a special
election become necessary because of the resignation of an SC member mid-year,
this year's Officers are responsible for running the special election, so 
any Officers should be at least tentatively available for that.

### Alternate Election Officer

In addition to the three Officers, the Elections Subproject will recommend one
Alternate.  This Alternate is available in case one of the Officers is unable
to complete the election, or is unavailable for post-election duties (such as
a special election).  Any current Officer may activate the Alternate if 
any Officer resigns or is unavailable.

The Alternate will be added to the Election Officers slack, but will have no
duties and will not participate in votes unless activated.

### K8s Infra Liaison

Currently election software runs on the Kubernetes cluster owned by the k8s-infra
team.  As such, the Election Officers may need troubleshooting and support from
k8s-infra in case of unexpected problems with the Elekto deployment or changes 
needed to the software.

As such, before the election starts, the Officers should reach out to k8s-infra
team and request one person who will be available to assist.  This person should
be available during most of the election period.  They must have the ability
to approve/modify services running on k8s.io.

If one of the Officers has these permissions, they may also serve in this role.

### Contributor-Comms Liaison

A big part of the election effort is making sure that voters and candidates
are aware of the election and kept up to date with constant reminders.  As such,
the Officers work directly with Contributor-Comms to send out a stream of
messages to the community.  

Well before the election starts, the Officers should reach out to contributor-comms
and ask them to assign one team member to handle election communications.  
This Comms member needs to have the ability to approve tweets.

If one of the Officers is a member of Contributor-Comms, they may double up
in this role.

[Elections Subproject]: /elections/README.md
