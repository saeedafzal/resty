VERSION := $(shell cat VERSION)
LD_FLAGS := -X main.version=$(VERSION)

build:
	go build -ldflags="$(LD_FLAGS) -s -w" -o bin/resty

run: build
	./bin/resty
