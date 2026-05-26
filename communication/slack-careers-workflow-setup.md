# Kubernetes Careers Slack Workflow Setup

This document details the technical configuration and how the job posting and resume workflows were built within the Kubernetes Slack workspace. Other administrators can use these specifications to maintain, audit, or recreate the system.

## System requirements and environment constraints

Before rebuilding or modifying this layout, verify the following workspace configurations and constraints:

* **Subscription and Branching Limitations**: The workspace runs on a Slack plan that does not support conditional branching within a single workflow file. 
* **Architecture Solution**: To work around this limitation, the system implements two independent, parallel parent workflows triggered by separate buttons instead of one workflow.
* **Featured Element Requirement**: Both user-facing workflows must be explicitly configured as "Featured Workflows" in Slack. This setting ensures the trigger buttons remain permanently visible and pinned directly to the channel message bar rather than hidden inside the conversation menu.
* **Target Channels**:
  * `#kubernetes-careers` (Public landing channel for user submissions)
  * `#triage-team` (Private administrative channel for triage team moderation)

## How submissions are processed

The workflow processes submissions through a defined sequence of states:

* **Submission capture**: A user submits data via public forms linked to the `#kubernetes-careers` message bar buttons.
* **Staging**: The submission is sent to the private `#triage-team` channel for validation.
* **Triage options**: 
  * **Approve**: Sends a confirmation message to the user and posts the approved job directly to the production `#kubernetes-careers` channel.
  * **Need more information**: Triggers an administrative sub-form from the `#triage-team` channel. This sends a direct message to the submitter containing the clarification request and an interactive response button. The submitter clicks the button to enter updates, which route directly back to the `#triage-team` channel.
  * **Reject**: Opens a rejection reason form, updates the private moderator logs, and routes policy feedback directly to the submitter.

## Workflow Step-by-Step Configuration

### New Job Opening

This workflow collects job details from users and sends them to moderators.

#### Workflow Initiation (Trigger)
* **Trigger Type**: Select `Starts from a link in Slack`.
* **Deployment Shortcut**: Navigate to the `#kubernetes-careers` conversation header, open the workflow properties, and select `Feature Workflow`. This pins the `New Job Opening` action button directly to the channel's message bar.

#### Form Data Collection
Configure the `Collect info in a form` step. Name the form `Job Submission Form` and define the following fields exactly as specified:
* `Job Title` (Short text)
* `Company Name` (Short text)
* `Work Arrangement` (Short text)
* `Full Description` (Long text)
* `Official Job Link` (Short text / URL)
* `Business Email` (Short text)
* `LinkedIn / GitHub Profile` (Short text / URL)
* `Salary Range` (Short text)
* `Location` (Short text)
* `Additional Details` (Long text)

#### Immediate Submitter Confirmation
Configure a `Send an 'only visible to you' message` step to provide immediate automated feedback to the user.
* **Recipient**: `{Person who used this workflow}`
* **Destination**: `{Channel where the workflow was used}`
* **Message Body Text**:
  > "Your submission has been received and is currently under review. You will receive a notification here once your post is live. If your post is not live within 24 hours, you can reach out to the `#slack-admins` channel."

#### Routing Data to Triage Team
Configure a `Send a message to` step to send the submission directly to the moderators.
* **Destination Channel**: `#triage-team`
* **Message Layout and Variable Mapping**:
  * **Header**: NEW SUBMISSION FOR REVIEW
  * **Submitter**: `{Person who submitted the form}` | `{Answer to: LinkedIn / GitHub Profile}`
  * **Business Email**: `{Answer to: Business Email}`
  * **Company**: `{Answer to: Company Name}`
  * **Role**: `{Answer to: Job Title}`
  * **Job Link**: `{Answer to: Official Job Link}`
  * **Location/Salary**: `{Answer to: Location}` | `{Answer to: Salary Range}`
  * **Description**: `{Answer to: Full Description}`
  * **Additional Details**: `{Answer to: Additional Details}`
* **Interactive Elements**: The message includes an `Approve` action button by default. Include two custom workflow buttons next to it labeled `Need More Information` and `Reject`.

#### Approve
Clicking the default `Approve` button triggers the following actions:
* Configure a `Send an 'only visible to you' message` step targeting `{Person who submitted the form}` inside `{Channel where the workflow was used}` containing:
  > "Your job posting for {Answer to: Job Title} is live."
* Configure a `Send a message to` step to post the final job to `#kubernetes-careers` using this format:
  * **Role**: `{Answer to: Job Title}`
  * **Company**: `{Answer to: Company Name}`
  * **Workplace**: `{Answer to: Work Arrangement}` | `{Answer to: Location}`
  * **Salary**: `{Answer to: Salary Range}`
  * **Description**: `{Answer to: Full Description}`
  * **Apply**: `{Answer to: Official Job Link}`
  * **Posted by**: `{Person who submitted the form}`

#### Need More Information
Clicking the custom `Need More Information` button requests additional information from the submitter:
* Configure a `Collect info in a form` step named `Need More Information Form` with the tracking fields:
  * `What details are missing from this submission?` (Long text)
  * `Requesting Info From:` (User selection menu)
* Configure a `Send a message to` step targeting `#triage-team` to update the moderator log:
  * **Header**: Awaiting Additional Information
  * **Issue**: `{Answer to: What details are missing from this submission?}`
  * **Submitter**: `{Answer to: Requesting Info From?}`
  * **Moderator**: `{Person who submitted the form}`
* Configure a `Send a message to` step to send a direct message to the submitter specified in `{Answer to: Requesting Info From:}` containing:
  > "Your submission requires the following details before it can go live:
  > 
  > {Answer to: What details are missing from this submission?}
  > 
  > Please click the button below to provide the requested information.
  > 
  > Thanks, 
  > Career Triage Team"
* Include an interactive button underneath the direct message block labeled `Provide details`. When the submitter selects this button, it opens a form titled `Additional Information Form` containing `Details` (Long text).
* Configure a final `Send a message to` step to send the submitter's response back to `#triage-team`:
  * **Header**: Submission Update
  * **Submitter**: `{Person who submitted the form}`
  * **Details**: `{Answer to: Details}`

#### Reject
Clicking the custom `Reject` button records the rejection and notifies the submitter:
* Configure a `Collect info in a form` step named `Rejection Form` with the fields:
  * `Reason for Rejection` (Long text)
  * `Submitter` (User selection menu)
* Configure a `Send a message to` step targeting `#triage-team` to log the decision:
  * **Header**: Reason for Rejection:
  * **Details (Quote Block)**: `{Answer to: Reason for Rejection}`
  * **Submitter**: `{Answer to: Submitter}`
  * **Moderator**: `{Person who submitted the form}`
* Configure a `Send a message to` send a direct message to the submitter specified in `{Answer to: Submitter}` containing:
  > "Hello,
  > Your submission has been reviewed and rejected for policy violations.
  > Moderator Feedback:
  > 
  > {Answer to: Reason for Rejection}
  > 
  > Next Steps: Please review the Job & Resume Posting Vetting Policy to ensure your submission meets all safety and formatting criteria. You may resubmit once these requirements are met."

### Resume Submission

This workflow collects candidate profiles from users and sends them to moderators.

#### Workflow Initiation (Trigger)
* **Trigger Type**: Select `Starts from a link in Slack`.
* **Deployment Shortcut**: Navigate to the `#kubernetes-careers` conversation header, open the workflow properties, and select `Feature Workflow`. This pins the `Resume Submission` action button directly to the message bar interface next to the job creation button.

#### Form Data Collection
Configure the `Collect info in a form` step. Name the form `Resume Form` and define the following fields exactly as specified:
* `Full Name` (Short text)
* `Resume Link` (Short text / URL)
* `Description` (Long text)

#### Immediate Submitter Confirmation
Configure a `Send an 'only visible to you' message` step to provide immediate automated feedback to the user.
* **Recipient**: `{Person who submitted the form}`
* **Destination**: `{Channel where the workflow was used}`
* **Message Body Text**:
  > "Your submission has been received and is currently under review. You will receive a notification here once your post is live. If your post is not live within 24 hours, you can reach out to the `#slack-admins` channel."

#### Routing Data to Triage Team
Configure a `Send a message to` step to send the resume details directly to the moderators.
* **Destination Channel**: `#triage-team`
* **Message Layout and Variable Mapping**:
  * **Name**: `{Answer to: Full Name}`
  * **Resume Link**: `{Answer to: Resume Link}`
  * **Details**: `{Answer to: Description}`
  * **Submitted by**: `{Person who submitted the form}`
* **Interactive Elements**: The message includes an `Approve` action button by default. Include two custom workflow buttons next to it labeled `Need More Information` and `Reject`.

#### Approve
Clicking the default `Approve` button triggers the following actions:
* Configure a `Send an 'only visible to you' message` step targeting `{Person who submitted the form}` inside `{Channel where the workflow was used}` containing:
  > "Your profile for {Answer to: Resume Link} is live."
* Configure a `Send a message to` step to post the resume details to `#kubernetes-careers`:
  * **Details (Quote Block)**: `{Answer to: Description}`
  * **Link**: `{Answer to: Resume Link}`
  * **Posted by**: `{Person who submitted the form}`

#### Need More Information
Follows the same steps described in the Need More Information section above.

#### Reject
Follows the same steps described in the Reject section above.
