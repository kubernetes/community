# Setting up the CNCF CLA check

If you are trying to sign the CLA so your PR's can be merged, please [read the
CLA docs](https://git.k8s.io/community/CLA.md)

If you are a Kubernetes GitHub organization or repo owner, and would like to
setup the Linux Foundation CNCF CLA check for your repositories, please read on.

## Setup the webhook

1. Go to the settings for your organization or webhook, and choose Webhooks from
   the menu, then "Add webhook"
    - Payload URL:
      `https://identity.linuxfoundation.org/lfcla/github/postreceive?group=284&comment=no&target=https://identity.linuxfoundation.org/projects/cncf`
      - `group=284` specifies the ID of the CNCF project authorized committers
        group in our CLA system.
      - `comment=no` specifies that our system should not post help comments
        into the pull request (since the Kubernetes mungebot does this).
      - `target=https://identity.linuxfoundation.org/projects/cncf` specifies
        what will be used for the "Details" link in GitHub for this status
        check.
    - Content Type: 'application/json'
    - Secret: Please contact [@idvoretskyi](mailto:ihor@cncf.io), and
      [@caniszczyk](mailto:caniszczyk@linuxfoundation.org).
    - Events: Let me select individual events
      - Push: **unchecked**
      - Pull request: checked
      - Issue comment: checked
      - Active: checked
1. Add the [@thelinuxfoundation](https://github.com/thelinuxfoundation) GitHub
user as an **Owner** to your organization or repo to ensure the CLA status can
be applied on PR's
1. After you send an invite, contact the [Linux
Foundation](mailto:helpdesk@rt.linuxfoundation.org); and cc [Chris
Aniszczyk](mailto:caniszczyk@linuxfoundation.org), [Ihor
Dvoretskyi](mailto:ihor@cncf.io), [Eric Searcy](mailto:eric@linuxfoundation.org)
(to ensure that the invite gets accepted).
1. Finally, open up a test PR to check that:
    1. webhooks are delivered correctly, which can be monitored in the
      “settings” for your org
    1. the PR gets the cla/linuxfoundation status

## Branch protection

It is recommended that the Linux Foundation CLA check be added as a strict
requirement for any change to be accepted to the master branch.

To do this manually:

1. Go to the Settings for the repository, and choose Branches from the menu.
1. Under Protected Branches, choose "master".
1. Check "Protect this branch".
1. Check "Require status checks to pass before merging", "Require branches to be
up to date before merging", and the "cla/linuxfoundation" status check.

Given the Kubernetes projects anticipates having "human reviewed" CLA
acceptance, you may not do the last step, but it is still recommended to enable
branch protection to require all changes to be done through pull requests,
instead of direct pushing that will never kick off a CLA check.

## Label automation

The label automation is done using the [CLA plugin in
prow](https://git.k8s.io/test-infra/prow/plugins/cla).  In order to turn on the
CLA labels on your repo, add it as appropriate within the
[plugins.yaml](https://git.k8s.io/test-infra/prow/plugins.yaml), and add the cla
plugin to it.

You also need to add [@k8s-ci-robot](https://github.com/k8s-ci-robot) as one of
the owners in the same org/repo, to ensure that it can add labels `cncf-cla:
yes` and `cncf-cla: no` based on the status published by the Linux Foundation
webhook.

The label automation may not be essential for your repository, if you’re not
using merge automation. For repos with maintainers doing manual merges, GitHub
protected branches may suffice.
