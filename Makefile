IMAGE_NAME=golang:1.22
CONTAINER_ENGINE?=$(shell command -v docker 2>/dev/null || command -v podman 2>/dev/null)

default: \
	generate \

reset-docs:
	git checkout HEAD -- ./sig-list.md ./sig-*/README.md ./wg-*/README.md

generate:
	go run ./generator/app.go

generate-containerized:
	$(CONTAINER_ENGINE) run --rm -e WHAT -v $(shell pwd):/go/src/app $(IMAGE_NAME) make -C /go/src/app generate

verify:
	@hack/verify.sh

verify-containerized:
	$(CONTAINER_ENGINE) run --rm -v $(shell pwd):/go/src/app $(IMAGE_NAME) make -C /go/src/app verify

test:
	go test -v ./generator/...

test-containerized:
	$(CONTAINER_ENGINE) run --rm -v $(shell pwd):/go/src/app $(IMAGE_NAME) make -C /go/src/app test

.PHONY: default reset-docs generate generate-containerized verify verify-containerized test test-containerized
