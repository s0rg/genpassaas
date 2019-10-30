BIN=bin/genpass.bin
CMD=./cmd/genpass
COVER=test.cover

GIT_HASH=`git rev-parse --short HEAD`
BUILD_DATE=`date +%FT%T%z`

LDFLAGS=-X main.GitHash=${GIT_HASH} -X main.BuildDate=${BUILD_DATE}
LDFLAGS_REL=-w -s ${LDFLAGS}

.PHONY: clean build

build: vet
	go build -ldflags "${LDFLAGS}" -o "${BIN}" "${CMD}"

release: vet
	go build -ldflags "${LDFLAGS_REL}" -o "${BIN}" "${CMD}"

docker: vet
	CGO_ENABLED=0 go build -ldflags "${LDFLAGS_REL}" -o "${BIN}" "${CMD}"
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
