Please read before using the data. 

| Data | Info |
| --- | --- |
Title | Kubernetes Contributor Experience Survey 2018
Authors | @parispittman, @jberkus, and many contributor experience members
Tool Used | SurveyMonkey; @idvoretskyi entered into the tool from the CNCF account and exported the data
Start | September 08, 2018 (soft launch on Slack); September 11, 2018 (full launch on kubernetes-sig-contribex@googlegroups.com)
End | October 1, 2018
Subject(s) | automation, community meeting, mentoring, communication, demographic information about contributors, events, 
Langauage | English
Data Processing | All personal identifiers have been removed. 73 respondents provided their email addresses for follow up. They have been scrubbed. 
Format | .csv
File Name | contribex-survey-2018.csv

Many column headers have been changed due to length, provide context, or they produced a two header column from ordinal scale and ranking questions. Changes will be documented below.

Some values represent a range of feelings/opinions. Check the question to find out the descriptive range. (ex: 1=least useful, 5=most useful)


### Two header column changes:

Columns: K-S  
Question: Please rate the below parts of the contribution process by how challenging they are, from 1 (not a problem) to 5 (a frequent blocker)

Columns: T-X  
Question: Which of the following tooling do you find useful?

Columns: Y, AA, AN, AX, BG, BO  
Question: open ended questions; may need to assign values to capture trends

Columns: AD-AM  
Question: What areas of Kubernetes do you contribute to? Please check all that apply.

Columns: AO-AW  
Question: What conferences have you previously attended or are planning to attend?

Columns: AZ-BF  
Question: How useful do you find each section of the Thursday's Community Meeting? (1 least useful; 5 most useful)

Columns: BH-BN  
Question: Some of the major projects we are working on are listed below, check one that is most important to you that we carry through to completion.

Columns: BP-BX  
Question: Of our various communications channels, please rate which ones you use and/or check most frequently on a 1-5 scale, where 1 is “never”, 3 is “several times a month” and 5 is “every day”.

Columns: BY-CG  
Question: Which of these channels is most likely to reach you first for news about decisions, changes, additions, and/or announcements to the contributor process or community matters?

### Full Question List:

1. How long have you been contributing to Kubernetes?
2. What level of the Contributor Ladder do you consider yourself to be on?
3. Are you interested in advancing to the next level of the Contributor Ladder?
4. What region of the world are you in?
5. What timezone are you most often in? (Check your UTC offset here)
6. How many other open source projects not in the Kubernetes ecosystem do you contribute to?
7. Please rate the below parts of the contribution process by how challenging they are, from 1 (not a problem) to 5 (a frequent blocker):  
Code/Documentation review  
Communication  
GitHub tools and processes (not our customized tooling)  
Finding the right SIG for your contributions  
Our CI, labels, and crafted customized automation  
Debugging test failures  
Finding appropriate issues to work on  
Setting up development environment  
Having PRs rejected  
8. Which of the following tooling do you find useful?  
automatic /retest of flakes (fejta-bot)  
automatic labeling of stale issues (fejta-bot)  
issue commands like /assign, /kind bug (k8s-ci-robot)  
PR commands like /approve, /lint (k8s-ci-robot)  
automatic merging of approved PRs (k8s-merge-robot and k8s-ci-bot)  
9. What tool above is the least useful and why? Wish something was automated that isn’t? List it here.
10. How do you perceive the current notification volume and utility?
11. Which areas could use additional automation?
12. Does your employer support your contributions to Kubernetes?
13. How often do you contribute upstream (code, docs, issue triage, etc.)?
14. What areas of Kubernetes do you contribute to? Please check all that apply.   
Core code inside of kubernetes/kubernetes  
Code inside of another repo in the kubernetes/* GitHub organization
Documentation  
Testing & Infrastructure  
Advocacy and events  
Community & Project management; SIG Chair etc.  
Plugins & Drivers (CSI, CNI, cloud providers)  
Related projects (Kubeadm, Helm, container runtimes, etc.)  
Don’t contribute yet, hoping to start soon  
15. Are there specific ways the project could make contributing easier for you?
16. What conferences have you previously attended or are planning to attend?  
Kubecon Europe 2017  
Kubecon North America 2017  
Kubecon Europe 2018  
Kubecon China 2018  
Kubecon North America 2018  
Kubecon Europe 2019  
Ecosystem events, eg. Helm Summit  
Other conferences with a Kubernetes track (like DockerCon or ContainerDay)
None
17. Do you have any suggestions on how to make the Contributor Summits more valuable to you (N/A if not applicable)?
18. How many Kubernetes Contributor Summits have you attended?
19. How useful do you find each section of the Thursday's Community Meeting? (1 least useful; 5 most useful)  
Demo  
KEP of the Week  
Devstats Chart of the Week  
Release Updates  
SIG Updates  
Announcements  
Shoutouts  
20. Any feedback on how the community meeting can be better?
21. Some of the major projects we are working on are listed below, check one that is most important to you that we carry through to completion:  
Mentoring programs for all levels  
GitHub Management  
Delivering valuable contributor summits at relevant events  
Launching a contributor site for a one stop shop for tailored project news, info, docs, and calendar  
Discovery and planning around communication and collaboration platforms to lead to potential centralization and/or consolidation  
Improving DevStats  
Keeping our community safe on our various communication platforms through moderation guidelines and new approaches  
22. What is missing from that list entirely? Why?
23. Of our various communications channels, please rate which ones you use and/or check most frequently on a 1-5 scale, where 1 is “never”, 3 is “several times a month” and 5 is “every day”.  
Google Groups/Mailing Lists  
Slack  
discuss.kubernetes.io  
Zoom video conferencing/meetings  
Discussions on Github Issues and PRs  
Unofficial channels (IRC, Hangouts, Twitter, etc.)  
StackOverflow  
YouTube recordings (community meetings, SIG/WG meetings, etc.)  
Google Docs/Forms/Sheets, etc (meeting agendas, etc)  
24. Which of these channels is most likely to reach you first for news about decisions, changes, additions, and/or announcements to the contributor process or community matters?  
kubernetes-dev mailing list  
Dedicated discuss.k8s.io forum for contributors  
Contributor Experience mailing list  
Slack  
Twitter  
A dedicated contributor site  
Kubernetes blog  
k/community repo in GH (Issues and/or PRs)
25. Do you think Slack adds value to the project for users and/or contributors?
26. Have you ever used the Help Wanted and/or Good First Issue labels on issues you file to find contributors?
27. Are you interested in mentoring a Kubernetes upstream Intern for Outreachy or Google Summer of Code? We are also looking for organizations to sponsor if your employer is interested.
28. Have you watched or participated in an episode of our YouTube mentoring series Meet Our Contributors?
29. How useful did you find Meet Our Contributors? (1 - not useful at all; 5 - extremely useful)If you have suggestions on improvements, leave those in the feedback box at the end of the survey.
30. What remains a blocker to becoming a mentor?
31. Would you like us to follow up with you about any of your answers, above? If so, share your email address here:
32. Do you have any comments, questions, or clarifications for your answers on this survey? Leave the general feedback here: