BIN_API=bin/genpass-api
CMD_API=./cmd/genpass-api

BIN_CLI=bin/genpass
CMD_CLI=./cmd/genpass-cli

COVER=test.cover

GIT_HASH=`git rev-parse --short HEAD`
BUILD_DATE=`date +%FT%T%z`

VER_PKG=github.com/s0rg/genpassaas/pkg/config

LDFLAGS=-X ${VER_PKG}.GitHash=${GIT_HASH} -X ${VER_PKG}.BuildDate=${BUILD_DATE}
LDFLAGS_REL=-w -s ${LDFLAGS}

.PHONY: clean cli

cli: vet lint
	go build -ldflags "${LDFLAGS_REL}" -o "${BIN_CLI}" "${CMD_CLI}"

api: vet lint
	go build -ldflags "${LDFLAGS_REL}" -o "${BIN_API}" "${CMD_API}"

docker: vet lint
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS_REL}" -o "${BIN_API}" "${CMD_API}"
	docker build -t s0rg/genpassaas:latest --no-cache=true .

vet:
	go vet ./...

test:
	go test -race -count 1 -v -coverprofile="${COVER}" ./...

test-cover: test
	go tool cover -func="${COVER}"

lint:
	golangci-lint run

clean:
	[ -f "${BIN_API}" ] && rm "${BIN_API}"
	[ -f "${BIN_CLI}" ] && rm "${BIN_CLI}"
	[ -f "${COVER}" ] && rm "${COVER}"
