BIN=bin/genpass.bin
CMD=./cmd/genpass

BIN_CLI=bin/genpass
CMD_CLI=./cmd/genpass-cli

COVER=test.cover

GIT_HASH=`git rev-parse --short HEAD`
BUILD_DATE=`date +%FT%T%z`

LDFLAGS=-X main.GitHash=${GIT_HASH} -X main.BuildDate=${BUILD_DATE}
LDFLAGS_REL=-w -s ${LDFLAGS}

.PHONY: clean cli build release

cli: vet
	go build -ldflags "${LDFLAGS_REL}" -o "${BIN_CLI}" "${CMD_CLI}"

build: vet
	go build -ldflags "${LDFLAGS}" -o "${BIN}" "${CMD}"

release: vet
	go build -ldflags "${LDFLAGS_REL}" -o "${BIN}" "${CMD}"

docker: vet
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS_REL}" -o "${BIN}" "${CMD}"
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
	[ -f "${BIN}" ] && rm "${BIN}"
	[ -f "${COVER}" ] && rm "${COVER}"
	[ -f "${BIN_CLI}" ] && rm "${BIN_CLI}"
