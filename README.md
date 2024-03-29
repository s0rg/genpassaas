[![Build](https://github.com/s0rg/genpassaas/workflows/ci/badge.svg)](https://github.com/s0rg/genpassaas/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/genpassaas)](https://goreportcard.com/report/github.com/s0rg/genpassaas)
[![Maintainability](https://api.codeclimate.com/v1/badges/56ca218b8a8d6940427a/maintainability)](https://codeclimate.com/github/s0rg/genpassaas/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/56ca218b8a8d6940427a/test_coverage)](https://codeclimate.com/github/s0rg/genpassaas/test_coverage)

[![License](https://img.shields.io/badge/license-MIT%20License-blue.svg)](https://github.com/s0rg/genpassaas/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/genpassaas)](go.mod)

# Generate Password as a Service

Microservice + cli example for password generation.

# features

- simple and clean design - `api` and `cli` share only password generation code, constants and version information
- version (git and build date) information stored in binaries by Makefile
- simplest Dockerfile - based on `scratch` image only
- zero-dependencies
- 95%+ test-covered codebase
- `api` follows [12-factor](https://12factor.net/) methodology
- `api` handlers take care of client `Accept` header value
- feature-complete `cli` version

## microservice

Exposes two endpoints, both can be called with GET and two params:
- len - generated passwords length (default = 16, min = 6, max = 32).
- num - number of passwords to generate (default = 10, min = 1, max = 64).

The endpoints are:
- `/v1/simple` - shuffles alphabet and randomly picks chars from it. Passwords, generated with this endpoint is pretty strong, but hard to remember.
- `/v1/smart` - generates prononceable (thereby easy to remember) passwords by repeating vowels and consonants with patterns.

By default both endpoints answer is `text/plain` block with `num` passwords separated by single LF (`\n`) byte.
This behavior can be changed via `Accept` header, if you set it to `application/json`, you will get json array of passwords.

### environment

- `KEY` set it to some random string, clients must provide same value via `Authorization: Bearer {KEY}` header
- `ADDR` full (interface:port) address to bind to (default = `localhost:8080`)

### building

- `make api` to get binary for manual execution
- `make docker` to get docker image with service
- `docker run -p 1337:8080 -e "KEY=my-app-key" s0rg/genpassaas:latest` to start container

### usage

- `curl -H "Authorization: Bearer my-app-key" "http://localhost:1337/v1/simple?len=10&num=10"`
- `curl -H "Authorization: Bearer my-app-key" "http://localhost:1337/v1/smart?len=10&num=10"`
- `curl -H "Accept: application/json" -H "Authorization: Bearer my-app-key" "http://localhost:1337/v1/smart?len=10&num=10"`

## cli

Provides binary to do all above.

### building

- `make cli`

### usage

```
Usage of genpass [flags]:

-count int
    count of passwords (default 10)
-gen string
    generator 'smart' or 'simple' (default "smart")
-len int
    length of each password (default 16)
-version
    show version and exit
```
