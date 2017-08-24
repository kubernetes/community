# The Contributor License Agreement

The [Cloud Native Computing Foundation][CNCF] defines the legal status of the
contributed code in a _Contributor License Agreement_ (CLA).

Only original source code from CLA signatories can be accepted into kubernetes.

This policy does not apply to [third_party] and [vendor].

## How do I sign?

#### 1. Read

  * [CLA for individuals] to sign up as an individual or as an employee of a signed organization.
  * [CLA for corporations] to sign as a corporation representative and manage signups from your organization.
  
#### 2. Sign in with GitHub.

Click
  * [Individual signup] to sign up as an individual or as an employee of a signed organization. 
  * [Corp signup] to sign as a corporation representative and manage signups from your organization.

Either signup form looks like this:

![CNCFCLA1](http://i.imgur.com/tEk2x3j.png)

#### 3. Enter the correct E-mail address to validate!

The address entered in the form must meet two constraints:
 
 * It __must match__ your  [git email] (the output of `git config user.email`)
   or your PRs will not be approved!

 * It must be your official `person@organization.com` address if you signed up
   as an employee of said organization.

![CNCFCLA2](http://i.imgur.com/t3WAtrz.png)

#### 4. Look for an email indicating successful signup.

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

If you have signup trouble, please explain your case on
the [CLA signing issue] and we (@sarahnovotny and @foxish), 
along with the [CNCF] will help sort it out.

Another option: ask for help at `helpdesk@rt.linuxfoundation.org`.

[CNCF]: https://www.cncf.io/community
[CLA signing issue]: https://github.com/kubernetes/kubernetes/issues/27796
[CLA for individuals]: https://github.com/cncf/cla/blob/master/individual-cla.pdf
[CLA for corporations]: https://github.com/cncf/cla/blob/master/corporate-cla.pdf
[Corp signup]: https://identity.linuxfoundation.org/node/285/organization-signup
[Individual signup]: https://identity.linuxfoundation.org/projects/cncf
[git email]: https://help.github.com/articles/setting-your-email-in-git
[third_party]: https://github.com/kubernetes/kubernetes/tree/master/third_party
[vendor]: https://github.com/kubernetes/kubernetes/tree/master/vendor
