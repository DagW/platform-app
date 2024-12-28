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

docker-test:
	docker build -t app/test:latest --target=build .
	docker run --rm docker.io/app/test:latest make test

docker-build:
	docker build \
	--target run \
	--build-arg buildTime=$(buildTime) \
  --build-arg revision=$(revision) \
  --build-arg version=$(version) \
  -t platform/app:latest .

clean:
	rm -rf bin

.PHONY: all install clean test build-docker
