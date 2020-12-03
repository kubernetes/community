# Requesting a Community Survey

Let us help you make your survey a success.

The Kubernetes project has access to the CNCF SurveyMonkey account for creating community surveys, and SIG-Contributor Experience includes people who can give advice on improving the quality of surveys, as well as promote them.  As such, what follows is the process for requesting such a community survey, in order to maximize its reach and data quality.

## What's a Community Survey?

Any survey requested by a Kubernetes SIG, WG, Team, or other standing Kubernetes community group (hereafter "Group") that targets some or all of the Kubernetes community and ecosystem is potentially a community survey.  As some examples, this would include:

* A survey by SIG-Cluster-Lifecycle on upgrade practices, targeted at Kubernetes users
* A survey by SIG-ContribEx on contributor events, targeted at current contributors
* A survey by the Release Engineering Team, targeted at vendors who repackage Kubernetes

Not included in community surveys are: surveys by specific vendors, surveys by other CNCF projects, surveys of target audiences outside the Kubernetes/CNCF community, or surveys that cannot be executed using SurveyMonkey and/or GoogleDocs.

If the survey you want to do is a community survey, then then next thing is to submit a request.

## Survey Request Process

### 1. Determine Goals and Content of Survey

The first step needs to happen in your community Group.  You need to determine:

1. Why you want to do a survey
2. Who you would like to answer it
3. What specific questions you want to ask

Ideally, answer those questions in that order; first decide what your goals are, next which audiences you're targeting, and finally compose the individual questions.

Your survey will also need a preface which explains to the audience what the survey is for, who will get to see its raw results, and where compiled data from it will be published.

### 2. Request Survey Review

Put your draft survey questions in a Google Doc, HackMD page, PR against one of your Group repos, or other format that supports comments by reviewers.  Then open an [issue](https://github.com/kubernetes/community/issues) against the community repo using the **Survey Request** template.

Members of SIG-ContribEx will then offer feedback on improving your survey, in the following areas:

* Are the questions and options clear and unambiguous?
* Are there obvious missing options to some questions?
* Is the preface clear and complete?
* Does the survey comply with the Kubernetes CoC and privacy practices?

Once you feel like the survey is as good as it's likely to be, we can publish it.

### 3. Request Survey Publication

By default, we use a SurveyMonkey account funded by the CNCF to publish surveys.  This is because SurveyMonkey offers a fairly rich set of survey options, and more importantly is accessible in China.  We can also publish surveys using the Kubernetes Google Drive account, if surveying Chinese developers or contributors is not a consideration.  GDrive makes sharing the survey results easier.

In either case, it will take a few days for a ContribEx volunteer to copy the questions and options into the survey tool, and then it can be ready to announce on your schedule.  Just comment on the issue giving your requested publication schedule, and mark it LGTM.

### 4. Contributor Communications Promotion (optional)

If you want your survey to reach a large number of contributors, developers, or users, then the Contributor Marketing subproject can help.  Use the [contributor comms issue template](https://github.com/kubernetes/community/issues/new/choose) to request help from that group.  They will ask you some questions about promoting the survey and assist you in getting the word out, via the Kubernetes social medial accounts, mailing lists, and blog, depending on the target audience and scope.

### 5. Collect Survey Results

Your survey can be open anywhere from a few days to a few months, depending on your needs.  In addition to any efforts Contributor-Comms may be making on your behalf, during this period your group should make sure to remind folks several times that the survey is open.  

On request (via the Github issue or Slack messages) ContribEx volunteers can supply intermediate information, such as the number of surveys completed or the current results data.

### 6. Publish Survey Results

When the survey is complete, the designated recipients for the raw data will receive a data dump in the form of a spreadsheet or CSV file, accompanied by copies of the simple graphs created by the survey system.  As this data may contain personally identifying information, the recipients are required to be a small group who will treat the data with care.

Depending on the availability of some ContribEx members, you may also be able to get more advanced visualizations and/or data analysis assistance.  This would be primarily correlating the answers to two or more questions, and providing graphs or charts for those correlations.  Request this in your issue and we'll follow-up to help figure out exactly what you need.

## Tips and Notes Around Surveys

Please note that we are subject to the GDPR and other privacy regulations.  As such, your survey will need to explain who will have access to the data and what it's for, and you need to adhere to those restrictions.  Among other things, this means that you can never share raw survey data on a public mailing list or or GitHub repo (summary data that contains no personal information should be fine).

Some additional tips for creating good surveys:

* Have goals: the overall survey, and each question, should be expected to provide useful data for your SIG.  If you can't figure out what you'd do with the response to a question, don't include that question.
* Ask more, simpler questions rather than fewer, complex ones.
* At the same time, try to produce the shortest survey that will actually give you the data you need.
* Remember to include "Don't Know" or "Not Applicable" options.
* Use "Other" options with care, as they tend to generate a lot of garbage input.
* Don't survey too often; survey fatigue will set in an you won't get responses.
