COMMIT     := $(shell git rev-parse HEAD)
VERSION    := $(shell git describe --tags $(COMMIT) 2> /dev/null || echo $(COMMIT))
BUILD_TIME := $(shell date +%FT%T%z)

LD_FLAGS   := -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.buildTime=$(BUILD_TIME)

all: build

build:
	go build -ldflags="$(LD_FLAGS) -s -w"

run: build
	./resty

upgrade:
	go get -u ./...
