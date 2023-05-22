current_time = $(shell date +"%Y-%m-%d:T%H:%M:%S")
linker_flags = '-s -X main.buildTime=${current_time}'

run:
	go run main.go

build:
	@echo "Building binaries..."
	go build -ldflags=${linker_flags}
