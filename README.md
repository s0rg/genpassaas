# Generate Password as a Service

 Microservice for password generation.

 Exposes two endpoints, both can be called with GET and two params:
 - len - generated passwords length (default = 8, min = 6, max = 64).
 - num - number of passwords to generate (default = 8, min = 1, max = 100).

 The endpoints are:
 - `/simple` - shuffles his alphabet and than randomly picks chars from it. Passwords, generated via this endpoint is pretty strong, but hard to remember.
 - `/smart` - tries to generate prononceable passwords (thereby easy to remember).

 By default both endpoints answer is `text/plain` block with `num` passwords separated by single LF (`\n`) byte.
 This behavior can be changed via `Accept` header, if you set it to `application/json`, 
 you will get json array of passwords.

# Environment

 - `KEY` set it to some random string, clients must provide same value via `Authorization: Bearer {KEY}` header
 - `ADDR` full (interface:port) address to bind to (default = `localhost:8080`)

# Building

 - `make build` to get binary for manual execution
 - `make docker` to get docker image with service
 - `docker run -p 1337:8080 -e "KEY=my-app-key" s0rg/genpassaas:latest` to start container

# Usage

```curl -H "Authorization: Bearer my-app-key" "http://localhost:1337/simple?len=10&cnt=10"```
```curl -H "Authorization: Bearer my-app-key" "http://localhost:1337/smart?len=10&cnt=10"```
```curl -H "Accept: application/json" -H "Authorization: Bearer my-app-key" "http://localhost:1337/smart?len=10&cnt=10"```

# License

MIT
