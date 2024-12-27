buildTime = $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
tag = $(shell git tag --points-at HEAD | grep "^vx" || echo "v0.0.0")
revision ?= $(shell git rev-parse --short HEAD)
version ?= $(tag)-$(revision)

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
