# ET

Email Tracker

Notify you when someone opens the email.

<!--
    I wrote this project just want to know whether someone had
    opened my email, waiting too long... She's getting married...
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

## Status

![Task Submitted](https://track.kfd.me/api/status/task/total?svg)
![Total Notified](https://track.kfd.me/api/status/notified/total?svg)
![Daily Notified](https://track.kfd.me/api/status/notified/daily?svg)

## Design

1. Users provide their **email address** to receive notifications.
2. Server generate an uniq `track ID` and returns to user,
    in the format of a 1x1 pixel png link.
3. User insert the png link into the email waiting to send.
4. When someone opens the link (`/t/xxxx-xxxx-xxxx`), server will send
    a notification email to the **email address** provided.
5. Since there is no way to identify the target's name, user can optionally
    set a target username or some comments to the `track task`.
6. Allow user to extend the notify times since the user could click
    the link by mistake.
7. Allow user to check the task status by the `track ID` in case
    of notify email failed to sent.

## API

- GET `/` Index page, provide a beautiful *task submit* portal.
- GET `/t/****` Track task handler, always returns a 1x1 pixel png file.
    Server will do something according to the track ID.
- `/api/` Raw API entrypoint, user can check task status, submit tasks
    and so on.
- `/api/task/`
  - POST `../submit` submit a new track task
  - POST `../resume?id=****` resume the stopped task
  - GET `../get?id=****` get task status, all notifications sent
- `/api/status`
  - GET `../task/total` *total* task handled
  - GET `../task/total?svg` return a status badge
  - GET `../notified/total?svg` return a status badge of *total* email sent
  - GET `../notified/daily?svg` return a status badge of *daily* email sent

## Constraints

- Same IP address can submit **10** tasks daily(per 24 hours).
    Can do this in memory, not a big deal.
- After sent **5** emails, automatically stop this task,
    user can resume the task once.
- Validate notifier's email, same email address can
    only receive **50** emails pre day.