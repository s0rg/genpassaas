BIN_API=bin/genpass-api
CMD_API=./cmd/api

BIN_CLI=bin/genpass
CMD_CLI=./cmd/cli

COVER=cover.out

GIT_TAG=`git describe --abbrev=0 2>/dev/null || echo -n "no-tag"`
GIT_HASH=`git rev-parse --short HEAD 2>/dev/null || echo -n "no-git"`
BUILD_AT=`date +%FT%T%z`

VER_PKG=github.com/s0rg/genpassaas/pkg/config
LDFLAGS=-X ${VER_PKG}.GitTag=${GIT_TAG} \
		-X ${VER_PKG}.GitHash=${GIT_HASH} \
		-X ${VER_PKG}.BuildDate=${BUILD_AT}
LDFLAGS_REL=-w -s ${LDFLAGS}

cli: test
	go build -ldflags "${LDFLAGS_REL}" -o "${BIN_CLI}" "${CMD_CLI}"

api: test
	go build -ldflags "${LDFLAGS_REL}" -o "${BIN_API}" "${CMD_API}"

docker: test
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS_REL}" -o "${BIN_API}" "${CMD_API}"
	docker build -t s0rg/genpassaas:latest --no-cache=true .

vet:
	go vet ./...

lint: vet
	golangci-lint run

test: lint
	go test -tags=test -race -count 1 -v -coverprofile="${COVER}" ./...

test-cover: test
	go tool cover -func="${COVER}"

clean:
	[ -f "${BIN_API}" ] && rm "${BIN_API}"
	[ -f "${BIN_CLI}" ] && rm "${BIN_CLI}"
	[ -f "${COVER}" ] && rm "${COVER}"
