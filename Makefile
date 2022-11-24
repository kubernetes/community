IMAGE_NAME=golang:1.18
export GOPROXY?=https://proxy.golang.org

default: \
	generate \

reset-docs:
	git checkout HEAD -- ./sig-list.md ./sig-*/README.md ./wg-*/README.md

generate:
	go run ./generator/app.go

generate-dockerized:
	docker run --rm -e WHAT -e GOPROXY -v $(shell pwd):/go/src/app:Z $(IMAGE_NAME) make -C /go/src/app generate

verify:
	@hack/verify.sh

verify-dockerized:
	docker run --rm -v $(shell pwd):/go/src/app:Z $(IMAGE_NAME) make -C /go/src/app verify

test:
	go test -v ./generator/...

test-dockerized:
	docker run --rm -v $(shell pwd):/go/src/app:Z $(IMAGE_NAME) make -C /go/src/app test

.PHONY: default reset-docs generate generate-dockerized verify test test-dockerized
