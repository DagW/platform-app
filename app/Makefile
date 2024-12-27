buildTime = $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
revision ?= $(shell git rev-parse --short HEAD)
version ?= $(shell git describe --tags --abbrev=7)

build:
	go build \
	-ldflags "-w -s -X main.buildTime=$(buildTime) -X main.revision=$(revision) -X main.version=$(version)" \
	-o ./bin/app \
	./cmd/main.go

run:
	go run \
	-ldflags "-w -s -X main.buildTime=$(date) -X main.revision=$(revision) -X main.version=$(version)" \
	./cmd/main.go

test:
	go test ./...

build-docker:
	docker build \
	--build-arg buildTime=$(buildTime) \
  --build-arg revision=$(revision) \
  --build-arg version=$(version) \
  -t nsight/app:latest .

clean:
	rm -rf bin

.PHONY: all install clean test build-docker
