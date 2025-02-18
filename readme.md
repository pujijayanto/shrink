## Overview

Shrink is an URL shortener. I build this project to practice while reading book [Let's Go](https://lets-go.alexedwards.net/).

## Features

- Shorten long URLs into compact, easy-to-share links.
- Track the number of times a shortened URL is accessed.
- Secure HTTPS connections using TLS.

## Prerequisites

- Go 1.18 or later
- PostgreSQL

## Getting Started


### Clone the Repository

```bash
git clone git@github.com:pujijayanto/shrink.git
cd shrink
```

### Enable SSL

You need to create a directory called `tls` in root project directory.

Find your Go standard library directory and call tool called `generate_cert.go`

If you’re using macOs followed the [official](https://go.dev/doc/install#install) install instructions,
then `generate_cert.go` file should be located under `/usr/local/go/src/crypto/tls.`

Then run

```bash
cd tls
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

### Start the application
Make sure you already enable SSL!.

Example you want to run the application with

- port 3003
- your database name is `shrink_dev`
- your postgres user of the database is `postgres`
- your postgres user password is `admin`

You can run by running

```bash
go run ./cmd/web -addr=":3003" -dsn="postgres://postgres:admin@localhost:5432/shrink_dev?sslmode=disable"
```

You can adjust the value based on your setup.

You can see help by running

```bash
go run ./cmd/web -help
```

And you can go to `https://localhost:3003`


### Running the test

Prerequisites:
1. prepare database for test
2. change the dns in [testutils_test.go](./internal/models/testutils_test.go)
3. clear cache `go clean -testcache`
4. run the test `go test ./...`

