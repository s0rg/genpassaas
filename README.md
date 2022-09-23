[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/genpassaas)](https://goreportcard.com/report/github.com/s0rg/genpassaas)

# Generate Password as a Service

Microservice + cli example for password generation.

## microservice

Exposes two endpoints, both can be called with GET and two params:
- len - generated passwords length (default = 16, min = 6, max = 32).
- num - number of passwords to generate (default = 10, min = 1, max = 64).

The endpoints are:
- `/v1/simple` - shuffles his alphabet and than randomly picks chars from it. Passwords, generated via this endpoint is pretty strong, but hard to remember.
- `/v1/smart` - tries to generate prononceable passwords (thereby easy to remember).

By default both endpoints answer is `text/plain` block with `num` passwords separated by single LF (`\n`) byte.
This behavior can be changed via `Accept` header, if you set it to `application/json`,
you will get json array of passwords.

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

# license

MIT
