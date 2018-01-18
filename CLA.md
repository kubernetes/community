# The Contributor License Agreement

The [Cloud Native Computing Foundation](https://www.cncf.io/community) defines
the legal status of the contributed code in a _Contributor License Agreement_
(CLA).

Only original source code from CLA signatories can be accepted into kubernetes.

This policy does not apply to [third_party](https://git.k8s.io/kubernetes/third_party)
and [vendor](https://git.k8s.io/kubernetes/vendor).

## What am I agreeing to?

There are two versions of the CLA:

1. One for [individual contributors](https://github.com/cncf/cla/blob/master/individual-cla.pdf)
submitting contributions on their own behalf.
1. One for [corporations](https://github.com/cncf/cla/blob/master/corporate-cla.pdf)
to sign for contributions submitted by their employees.

It is important to read and understand this legal agreement.

## How do I sign?

#### 1. Log into the Linux Foundation ID Portal with Github

Click one of:
  * [Individual signup](https://identity.linuxfoundation.org/projects/cncf) to
  sign up as an individual or as an employee of a signed organization.
  * [Corporation signup](https://identity.linuxfoundation.org/node/285/organization-signup)
  to sign as a corporation representative and manage signups from your organization.

Once you get to the sign in form, click "Log in with Github":

![CNCFCLA1](http://i.imgur.com/tEk2x3j.png)

#### 2. Create Linux Foundation ID Portal account with correct e-mail address

Ensure that the e-mail address you use when completing this form matches the one
you will use for your commits.

If you are signing up as an employee, you must use your official
person@organization.domain email address in the CNCF account registration page.

![CNCFCLA2](http://i.imgur.com/t3WAtrz.png)

#### 3. Complete signing process

Once you have created your account, follow the instructions to complete the
signing process via Hellosign.

#### 4. Ensure your Github e-mail address matches address used to sign CLA

Your Github email address __must match__ the same address you use when signing
the CLA. Github has [documentation](https://help.github.com/articles/setting-your-commit-email-address-on-github/)
on setting email addresses.

You must also set your [git e-mail](https://help.github.com/articles/setting-your-email-in-git)
to match this e-mail address as well.

If you've already submitted a PR you can correct your user.name and user.email
and then use use `git commit --ammend --reset-author` and then `git push` to
correct the PR.

#### 5. Look for an email indicating successful signup.

> The Linux Foundation
>
> Hello,
>
> You have signed CNCF Individual Contributor License Agreement.
> You can see your document anytime by clicking View on HelloSign.
>

Once you have this, the CLA authorizer bot will authorize your PRs.

![CNCFCLA3](http://i.imgur.com/C5ZsNN6.png)

## Troubleshooting

If you are having problems with signed the CLA send a mail to: `helpdesk@rt.linuxfoundation.org`.

Someone from the CNCF will respond to your ticket to help.

## Setting up the CNCF CLA check

If you are a Kubernetes GitHub organization or repo owner, and would like to setup
the Linux Foundation CNCF CLA check for your repositories, please 
[read the docs on setting up the CNCF CLA check](/setting-up-cla-check.md)
