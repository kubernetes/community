## Kubernetes Elections 

This document will outline how to conduct a Kubernetes Steering Committee Election. See the [Steering Committee Charter](https://github.com/kubernetes/steering/blob/master/charter.md) for more information of how the committee decides when to have an election, the method, and the maximal representation. 

## Steering Committee chooses Election Deadlines and Officers

- Steering Committee selects the Election Officers
- Dates should be in UTC time, use a [world clock service](https://www.timeanddate.com/worldclock/fixedtime.html?msg=Election+Test&iso=20181101T00&p1=%3A&ah=10) in documentation and email announcements so that end users see the correct time and date based on where they live. 


### SC Selects the following dates:

- Recommend the month of October to not collide with a release or end of a quarter.
- Nomination and Voter Registration period start
- Nomination period end (At least a two week period)
- Voter Registration Deadline
- Link to voter registration process (doesn’t exist yet)
- Election period start
  - It takes time to create the poll in CIVS, so don’t give a specific hour, instead say “Morning of the 10th” or something vague. 
- Election period stop
  - CIVS needs to be manually stopped, so an actual person needs to click for the poll to stop, so this needs to be a human friendly time. 
- Results announcement date

## Process

1. Election officers prepare the election repository
   - Make github.com/kubernetes/community/elections/$YEAR
   - Make github.com/kubernetes/community/elections/$YEAR/README.md, this is the voter’s guide. 
     - Copy over the voter’s guide from the previous year. The voter’s guide is the single source of truth for the election that year! All announcements and notices should link to this document. 
     - Update with new dates, candidates, and procedures (if necessary).
   - Announce to the candidates to submit PRs with their platform statement (if they desire), 300 word limit. Each platform document lives in the elections/$YEAR directory, with the voter’s guide (README.md) acting as the index. 

2. Announce voting schedule to community

- Should mostly be links to the voter guide and the Steering Committee Charter
- On kubernetes-dev list, kubernetes-dev slack, and twitter

3. Executing the Election in CIVS

- Use [CIVS](http://civs.cs.cornell.edu/civs_create.html) to create the election, which CIVS calls a poll. Once you send out the ballots you cannot UNSEND the emails, ensure everything in the form is correct!
- Name of the poll - “Kubernetes Steering Committee Election for $YEAR”
- Name of supervisor - “Kubernetes Steering Committee”
- Email - elections@kubernetes.io : Googlegroups doesn’t work here. This mail should resolve to members of the steering committee AND the election officers.
- Date and Time: Write in the date and time the election will stop. This field is not programmatic, the election is stopped by hand, so you can write this in plain text.
- Description: This election is to nominate the steering committee for the Kubernetes project. Select the three(3) candidates, by order of preference. Please see the voter's guide for more information.  PLEASE NOTE: "No opinion" is also a voting option if you do not feel comfortable ranking every single candidate.
- Add the candidate list to the form
- How many choices will win: This number needs to be set to the amount of open seats of a given election
- More options, check the boxes for:
  - Do not release results to all voters.
  - Enable detailed ballot reporting.
  - Allow voters to select “no opinion” for some choices.
- Click create poll, this will send elections@kubernetes.io an email with instructions.
- It will send you a link to “Poll Control”, bookmark this generated page as this is where you will add voters and also resend ballots to people if their ballot gets lost or filtered. 
- This page is where the “Start Poll” and “Stop Poll” buttons are, start the poll.
- Paste in the registered voters and click add voters.
  - It will mail the ballots to the participants.
  - It does duplicate detection so multiple entries are fine. 
- Leave the poll open for the duration of voting. 
  - Remember to send a 24 hour reminder before closing the poll. 

## Roles and Responsibilities:

### Steering Committee

- Select election dates
- Select Election Officers
- Select criteria for Kubernetes Members of Standing 
- Must refrain from endorsing or otherwise advocating for any candidate
- Is allowed to vote in the election
- Announces results of the election to the community
- Commit the results of the election to the Kubernetes Community repository

### Election Officers

- Must be a Kubernetes Member of Standing
- Cannot be running for office in the current election
- Cannot be a current member of the steering committee 
- Generates the voter guide
- Tracks candidates 
- Monitors kubernetes-dev for nominations
  - Keeps track of nominees in a spreadsheet
  - Ensures that each nominee has the required nominations from three different employers (as stated in the charter)
  - All nominations are conducted in the public, so sharing this sheet during the nomination process is encouraged
- Accepts/Reviews pull requests for the candidate platforms
  - The community generally assists in helping with PRs to give the candidates a quick response time
- Update the community regularly via the community meeting
- Post on behalf of the steering committee if necessary
- Posting deadlines and reminders to kubernetes-dev, twitter, and slack. 
- Reissues ballots from CIVS to voters who might have not received their ballot.
- Miscellaneous election related tasks as decided by the steering committee.
- Must refrain from endorsing or otherwise advocating for any candidate.
- Must refrain from discussing the election specifics during the election period. 
- Guard the privacy of the email addresses of voters
- It is impossible for the election officers to see the results of the election until the election ends; for purposes of transparency with the community it is encouraged to release some statistics during the election (ie. “65% of the community has voted so far!”) 
- Ensure that the election results are handed over to the steering committee. 
- Is allowed to vote in the election