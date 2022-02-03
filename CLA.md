# The Contributor License Agreement

The [Cloud Native Computing Foundation](https://www.cncf.io) (CNCF) defines
the legal status of the contributed code in two different types of _Contributor License Agreements_
(CLAs), [individual contributors](https://github.com/cncf/cla/blob/master/individual-cla.pdf) and [corporations](https://github.com/cncf/cla/blob/master/corporate-cla.pdf).

Kubernetes can only accept original source code from CLA signatories.

This policy does not apply to [third_party](https://git.k8s.io/kubernetes/third_party)
and [vendor](https://git.k8s.io/kubernetes/vendor).

It is important to read and understand this legal agreement.

## How do I sign?

After creating your first Pull Request the linux-foundation-easycla bot will respond with information regarding your CLA status along with a link to sign the CLA.

<img width="1065" alt="EasyCLA bot" src="https://user-images.githubusercontent.com/69111235/152226443-f6fe61ee-0e92-46c5-b6ea-c0deb718a585.png">

#### 1. Authorize EasyCLA to read some of your GitHub information

<img width="554" alt="GitHub EasyCLA Authorization" src="https://user-images.githubusercontent.com/69111235/152228712-7d22f9d0-9f3c-4226-9ee0-bacba4b47725.png">

Click on the "Please click here to be authorized" link to navigate to the GitHub Authorize Linux Foundation: EasyCLA page. Then click Authorize LF-Engineering to give Linux Foundation access to your GitHub account.

#### 2. Select from the two types of contributor

<img width="1407" alt="EasyCLA" src="https://user-images.githubusercontent.com/69111235/152224818-1246453a-b086-4a57-9d14-c10d62ad438f.png">


After authorizing EasyCLA, you will be redirected to a page to identify which type of contributor you are. 
Select either one:
  * Individual Contributor: To sign up as an individual or as an employee of a signed organization.
  * Corporate Contributor: To sign as a corporation representative and manage signups from your organization.

#### 3. Sign the CLA

Once you select your type of contributor proceed to Sign the CLA and follow the instructions to complete the signing process through DocuSign.

**Ensure your GitHub e-mail address matches e-mail address used to sign CLA**

After you have carefully filled up the information, Click "Finish" and you will be redirected to your GitHub Pull Request page.

#### 4. Look for an email indicating successful signup.

> Hello,
> 
> This is a notification email from EasyCLA regarding the project Cloud Native Computing > Foundation (CNCF).
> 
> The CLA has now been signed. You can download the signed CLA as a PDF here.
> 
> If you need help or have questions about EasyCLA, you can read the documentation or reach out to us for support.
> 
> Thanks,
> EasyCLA Support Team



#### 5. Validate your CLA

Once you are redirected to your GitHub Pull Request, enter `/easycla` to check whether your CLA was signed successfully or not.


## Changing your Affiliation

If you've changed employers and still contribute to Kubernetes, your affiliation
needs to be updated. The Cloud Native Computing Foundation uses [gitdm](https://github.com/cncf/gitdm)
to track who is contributing and from where. Create a pull request on the [gitdm](https://github.com/cncf/gitdm)
repository with a change to the corresponding developer affiliation text file.
Your entry should look similar to this:

```
Jorge O. Castro*: jorge!heptio.com, jorge!ubuntu.com, jorge.castro!gmail.com
Heptio
Canonical until 2017-03-31
```

## Troubleshooting

If you encounter any problems signing the CLA and need further assistance, log a ticket by clicking on the link [please submit a support request ticket](https://jira.linuxfoundation.org/plugins/servlet/theme/portal/4) from the EasyCLA bot's response. Someone from the CNCF will respond to your ticket to help.

Should you have any issues using the LF Support Site, send a message to the
backup e-mail support address <login-issues@jira.linuxfoundation.org>

## Setting up the CNCF CLA check

If you are a Kubernetes GitHub organization or repo owner and would like to setup
the Linux Foundation CNCF CLA check for your repositories, [read the docs on setting up the CNCF CLA check](/github-management/setting-up-cla-check.md)


[Linux Foundation Support Site]: https://support.linuxfoundation.org/
