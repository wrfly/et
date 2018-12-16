# ET

Email Tracker

Notify you when someone opens the email.

<!--
    I wrote this project just want to know whether someone had
    opened my email, waiting too long... And I like her.
    :(
-->

[![Go Report Card](https://goreportcard.com/badge/github.com/wrfly/et)](https://goreportcard.com/report/github.com/wrfly/et)
[![Master Build Status](https://travis-ci.org/wrfly/et.svg?branch=master)](https://travis-ci.org/wrfly/et)
[![GoDoc](https://godoc.org/github.com/wrfly/et?status.svg)](https://godoc.org/github.com/wrfly/et)
[![License](https://img.shields.io/github/license/wrfly/et.svg)](https://github.com/wrfly/et/blob/master/LICENSE)
[![Docker Pulls](https://img.shields.io/docker/pulls/wrfly/et.svg)](https://hub.docker.com/r/wrfly/et)
[![Image Size](https://img.shields.io/microbadger/image-size/wrfly/et.svg)](https://hub.docker.com/r/wrfly/et)
[![GitHub release](https://img.shields.io/github/release/wrfly/et.svg)](https://github.com/wrfly/et/releases)

---

![Tasks Submitted](https://track.kfd.me/api/status/tasks/total?svg)
![Total Notified](https://track.kfd.me/api/status/notified/total?svg)
![Daily Notified](https://track.kfd.me/api/status/notified/daily?svg)

## Design

1. Users provide their **notify-email address**.
2. Server generate an uniq `track ID` and returns to user,
    in the format of a 1x1 pixel png link.
3. User insert the png link into the email waiting to send.
4. When someone opens the link (`/t/xxxx-xxxx-xxxx`), server will send
    a notification email to the **notify-email address**, first 5 times.
5. Since there is no way to identify the target's name, user can optionally
    set a target username or some comments to the `track task`.
6. Allow user to extend the notify times since the user could click
    the link by mistake.
7. Allow user to check the task status by the `track ID` in case
    of notify email failed to sent.

## API

- `/` Index page, can provide a beautiful *task submit* portal.
    But can also handle request from `curl` (reuse the API handler).
- `/t/****` Track task handler, always returns a 1x1 pixel png file.
    Server will do something according to the track ID.
- `/api/` Raw API entrypoint, user can check task status, submit tasks
    and so on.
