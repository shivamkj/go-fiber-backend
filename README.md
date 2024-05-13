# Qnify Learning System Backend Server

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)

[Qnify](https://qnify.pages.dev/) is an open source modern customisable Learning Management System, with focus on Learning not on Management. It is designed to be **fast** with **cost** in mind.

## Commands

- Run in prod mode: `go run -tags prod main.go`
- Format schema SQL files: `sqlfluff format schema.sql --dialect postgres`
- Start Redis (without disk save): `redis-server --save "" --appendonly no`
- Generate Token Secret Keys: `openssl rand -hex 32`

### Copyright

Copyright (©) 2024 - Shivam Kumar Jha
