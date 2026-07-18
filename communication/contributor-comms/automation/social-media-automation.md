# Social Media Automation

Contributor Comms manages the social media accounts for the Kubernetes community.
We use automation to alleviate the workload of posting blog posts, LWKD, and CVEs
to social media. This document describes the automation process in place to
push the various posts out to the social media accounts for the project.

## Overview

New posts published to any of the project's RSS feeds are automatically queued
and published to six social media accounts. The pipeline is fully automated via
Zapier and Buffer — no manual posting is required.

## Pipeline

```mermaid
flowchart TB
    subgraph feeds["RSS Feeds"]
        direction LR
        f1["k8s.io"]
        f2["k8s.dev"]
        f3["etcd.io"]
        f4["LWKD"]
        f5["CVEs"]
    end

    subgraph zaps["Zapier — one zap per network"]
        direction LR
        z1["Mastodon zap"]
        z2["LinkedIn zap"]
        z3["K8s.io Twitter zap"]
        z4["K8s.io Bluesky zap"]
        z5["K8s Contrib Twitter zap"]
        z6["K8s Contrib Bluesky zap"]
    end

    subgraph buffer["Buffer queues"]
        direction LR
        b1["Mastodon queue"]
        b2["LinkedIn queue"]
        b3["K8s.io Twitter queue"]
        b4["K8s.io Bluesky queue"]
        b5["K8s Contrib Twitter queue"]
        b6["K8s Contrib Bluesky queue"]
    end

    subgraph social["Social media accounts"]
        direction LR
        s1["Mastodon"]
        s2["LinkedIn"]
        s3["K8s.io Twitter"]
        s4["K8s.io Bluesky"]
        s5["K8s Contrib Twitter"]
        s6["K8s Contrib Bluesky"]
    end

    feeds --> z1 & z2 & z3 & z4 & z5 & z6

    z1 --> b1 --> s1
    z2 --> b2 --> s2
    z3 --> b3 --> s3
    z4 --> b4 --> s4
    z5 --> b5 --> s5
    z6 --> b6 --> s6

    classDef plain fill:transparent,stroke:#888,stroke-width:1px,color:inherit
    class f1,f2,f3,z1,z2,z3,z4,z5,z6,b1,b2,b3,b4,b5,b6,s1,s2,s3,s4,s5,s6 plain
    style feeds fill:transparent,stroke:#888
    style zaps fill:transparent,stroke:#888
    style buffer fill:transparent,stroke:#888
    style social fill:transparent,stroke:#888
```

## Stages

### 1. RSS feeds

The pipeline is triggered by three RSS feeds:

- `k8s.io`
- `k8s.dev`
- `etcd.io`
- `LWKD`
- `CVEs`

A new item on any of these feeds is what kicks off the automation.

### 2. Zapier

There is one Zapier zap per social media network. There are 18 total zaps.
Blog RSS feeds are one group of zaps (6), LWKD makes another group of zaps
(6), and CVEs make the last group of zaps (6). Each zap watches all three
RSS feeds, and when a new item appears it builds a post from the item's
**Title** and **Link**, then adds that post to the network's Buffer queue.

### 3. Buffer

Each connected account has its own Buffer queue. Buffer holds the queued posts
and handles the actual scheduling and publishing to the account.

### 4. Social accounts

Posts are published to the following six accounts:

- Mastodon
- LinkedIn
- K8s.io Twitter
- K8s.io Bluesky
- K8s Contrib Twitter
- K8s Contrib Bluesky