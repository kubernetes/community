---
title: "Slack Announcements Moderation & Vetting Criteria"
description: |
  Guidelines and policies for posting announcements in the Kubernetes
  Slack #announcements channel.
---

<!-- overview -->

This document outlines the vetting guidelines used by Kubernetes Slack moderators to manage the `#announcements` channel, which uses a moderated queue. Submissions are reviewed in a private triage channel, where moderators can approve, request more information, or reject based on the criteria below. Approved announcements go to `#announcements`. Direct posting remains restricted to Slack admins.

<!-- body -->

## 1. Acceptance Criteria

Announcements is our largest-distribution channel, and as such posts to Announcements are expected to be infrequent, major news. They should not include regular reminders or follow-ups unless those convey critical information for global Kubernetes users. The Release Team generally makes only 2-4 posts total on Announcements per release cycle, including the announcement of the final release. Event announcements apply only to global community events, such as KubeCons, and not regional events or deadline reminders. Only deprecations with major user impact, such as dockershim and Ingress NGINX, go on Announcements.

A submission should fall under one of the following categories: Security Advisory, Release, Elections, Worldwide Community Event, Slack Admin Notice, or Deprecation / Retirement. It should also include a clear title and description, with a working link where applicable. Submissions are made through a Slack workflow (Title, Category, Description, Link), and typically come from the Security Response Committee, release teams, and Slack admins.

## 2. Rejection Flags

Submissions are rejected if they fall outside the scope defined above, such as content unrelated to Kubernetes, unsolicited promotional material, or announcements that don't meet the importance and frequency bar for this channel. If a submission is otherwise appropriate but missing a title, category, or description, moderators will request the missing details instead of rejecting it outright.

## 3. Feedback & Appeals

Submitters who are rejected or asked for more information receive a message explaining the reason. If someone believes a rejection was made by mistake, they may reach out to the `#slack-admins` channel for review.
