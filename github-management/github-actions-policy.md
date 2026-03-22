The purpose of this policy is to establish mandatory security requirements while using GitHub Actions in workflow files across all repositories under all Kubernetes github organizations.

**All GitHub Actions MUST be referenced using commit SHA hashes.**

```yaml
# REQUIRED - Pin to commit SHA
uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

# PROHIBITED - Mutable references
uses: actions/checkout@v4        # Tags can be force-pushed
uses: actions/checkout@main      # Branches move
uses: actions/checkout@latest    # Undefined reference
```

### Rationale

Mutable references (tags, branches, `latest`) can be force-updated to point to different commits. An attacker who compromises an action's repository can inject malicious code by modifying what these references point to, creating supply chain vulnerabilities.

Commit SHA hashes are cryptographically immutable and cannot be changed, preventing such attacks.

Recent incidents demonstrate the risks of mutable references:
- https://github.com/aquasecurity/trivy/security/advisories/GHSA-69fq-xp46-6x23
- Discussion: https://kubernetes.slack.com/archives/CD6LAC15M/p1774101470025069

Additional context:
- https://docs.github.com/en/actions/security-guides/security-hardening-for-github-actions#using-third-party-actions

## Requirements

1. All `uses:` statements in workflow files MUST reference actions using 40-character commit SHA hashes
2. New workflows MUST comply before merge
3. Existing workflows MUST be updated to comply
4. Repositories SHOULD enable Dependabot for GitHub Actions to automatically update SHA-pinned actions to newer versions
