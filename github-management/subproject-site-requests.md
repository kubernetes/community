# Subproject Site Requests

Official Kubernetes subprojects can request a domain and site hosting for
their project specific documentation. The Kubernetes community has standardized
on [Netlify] for this purpose. Netlify was chosen for its ease of use,
integrations such as offering automatic site previews per PR, and its support
for multiple site generation frameworks.

**NOTE:** This process is only for requesting a domain and hosting, not the
building and design of the site itself. The responsibility for maintaining the
subproject site belongs to the subproject itself.

If you are unsure of what site framework to use, [Hugo] is a lightweight Go
based framework. It's used for the Kubernetes website and several subprojects.


- [Requesting a Netlify Site](#requesting-a-netlify-site)
  - [Example Netlify Configuration](#example-netlify-configuration)
- [Requesting a Subproject Domain](#requesting-a-subproject-domain)
  - [Example Subproject Domain Request](#example-subproject-domain-request)
- [Processing a Netlify Site Request](#processing-a-netlify-site-request)


## Requesting a Netlify Site

- Update the project to include a [site specific][site-config] configuration
  file at the root of the repository ([`netlify.toml`][site-config]).
  [Below](#example-netlify-configuration) is a [Hugo] based example that you
  may use as a reference.

- Create an issue in the [kubernetes/org] repository using the
  [Netlify Site Request Template]. It will ask for:
  - Requesting sub-project and associated information. Linking to an issue or
    other supporting material for the Netlify request associated with the
    subproject is encouraged and will speed processing of the request.
  - Desired domain name. It should follow the pattern of
    `<sub-project>.sigs.k8s.io` and match the
    [request](#requesting-a-subproject-domain) for a subproject domain.

- A member of the GitHub Admin team handling Netlify requests will respond and
  follow up with any questions in the issue.

- Once Netlify has been configured, you can preview the site using the Netlify
  URL (example: `kubernetes-sigs-foo.netlify.com`). After this looks good to you,
  [request a subproject domain](#subproject-domain-request) and reference the
  issue you created.

Once complete, the site should be accessible.

### Example Netlify Configuration

```toml
[build]
base = "site/"
publish = "site/public/"
command = "hugo"

[build.environment]
HUGO_VERSION = "0.53"

[context.production.environment]
HUGO_ENV = "production"
HUGO_BASEURL = "https://foo.sigs.k8s.io/"

[context.deploy-preview]
command = "hugo --enableGitInfo --buildFuture -b $DEPLOY_PRIME_URL"

[context.branch-deploy]
command = "hugo --enableGitInfo --buildFuture -b $DEPLOY_PRIME_URL"
  ```

## Requesting a Subproject Domain

Subproject domains may be requested for a Kubernetes Org managed Netlify site in
the form of: `<project>.sigs.k8s.io`.

To do so, create an issue using the [DNS Update Request] issue template in the
[kubernetes/k8s.io] repository. Use the below guidelines for your
[DNS Update Request].

- **Type of DNS update:** Create
- **Domain being modified:** `k8s.io`
- **New DNS Record:**
  ```yaml
  # https://github.com/kubernetes-sigs/<project> (<contact or owners>)
  <project>.sigs:
    type: CNAME
    value: <netlify url>
  ```
- **Reason for update:** Provide a description and link to the Netlify request
  issue made in [kubernetes/org].


### Example Subproject Domain Request

````
### Type of DNS update:
Create

### Domain being modified:
k8s.io

### New DNS Record:
```yaml
# https://github.com/kubernetes-sigs/foo (@bar)
foo.sigs:
  type: CNAME
  value: kubernetes-sigs-foo.netlify.com
  ```
### Reason for update:
The sig-foo subproject has requested a netlify site to host its documentation. ref #0000, #0000.
````


## Processing a Netlify Site Request

**NOTE:** For use by the GitHub Admin team. You must be both an Org owner and
Netlify admin to follow the procedure.

- Login to [Netlify] and from the home dashboard select **New Site from Git**.

- On the _"Create a new site"_ page, select GitHub. It will then prompt you to
  authorize the application for the desired GitHub Organization. Select the
  GitHub Organization and the desired repo.

- In the _"Deploy Options"_ ensure the **Owner** is set to `Kubernetes Docs` and
  **Branch to deploy** is set to `master`. The _"Basic build settings"_ will be
  autopopulated with the values provided in `netlify.toml`. Deploy the site.
  It will take you to the _"Site overview"_ page.

- Navigate to the **Site Settings** and then change the Site name following the
  convention `kubernetes-sigs-<repo/project name>` e.g. `kubernetes-sigs-foo`.
  This will be used as both the Netlify site name and in the auto-generated PR
  based previews.

- From the left hand menu, select **Domain management**.

- Select **[Add custom domain]**. Then enter the domain name requested in the
  issue. It should follow the pattern of `<subproject-name>.sigs.k8s.io`.
  Note that HTTPS will not be enabled right away. Once the DNS configuration is
  updated, HTTPS will be enabled automatically.

- Follow up with the requestor in the issue and let them know the site has been
  deployed and give them the Netlify site url (site name + `netlify.com`).
  Example `kubernetes-sigs-foo.netlify.com`. They may use this address for
  testing and use before their DNS Request has been processed.

Once complete, the rest of the Netlify site configuration can be handled by the
project owner in their [netlify.toml][site-config] config.


[netlify]: https://www.netlify.com/
[kubernetes website]: https://git.k8s.io/website
[hugo]: https://gohugo.io
[dns update request]: https://github.com/kubernetes/k8s.io/issues/new/choose
[kubernetes/org]: https://git.k8s.io/org
[kubernetes/k8s.io]: https://git.k8s.io/k8s.io
[netlify site request template]: https://github.com/kubernetes/org/issues/new/choose
[dns zone config]: https://git.k8s.io/k8s.io/dns/zone-configs/k8s.io.yaml
[site-config]: https://www.netlify.com/docs/netlify-toml-reference/
[add custom domain]: https://www.netlify.com/docs/custom-domains/