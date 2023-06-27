COMMIT := $(shell git rev-parse HEAD)
VERSION := $(shell git describe --tags $(COMMIT) 2> /dev/null || echo $(COMMIT))
COMMIT := $(shell git rev-parse HEAD)
BUILD_TIME := $(shell date +%FT%T%z)
LD_FLAGS := -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.buildTime=$(BUILD_TIME)

run:
	go run -ldflags="$(LD_FLAGS)" main.go

build:
	go build -ldflags="$(LD_FLAGS)"
