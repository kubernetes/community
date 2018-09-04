IMAGE_NAME=golang:1.9

default: \
	generate \

reset-docs:
	git checkout HEAD -- ./sig-list.md ./sig-*/README.md ./wg-*/README.md

generate:
	go run ./generator/app.go

generate-dockerized:
	docker run --rm -e WHAT -v $(shell pwd):/go/src/app:Z $(IMAGE_NAME) make -C /go/src/app generate

generate-icons-drawio:
	cd icons && ./tools/draw.io/update.sh

generate-icons-drawio-dockerized:
	docker run --rm -v $(shell pwd):/build -w '/build' alpine sh -c 'cd icons && ./tools/draw.io/update.sh'

verify:
	@hack/verify.sh

test:
	go test -v ./generator/...

test-dockerized:
	docker run --rm -v $(shell pwd):/go/src/app:Z $(IMAGE_NAME) make -C /go/src/app test

.PHONY: default reset-docs generate generate-dockerized generate-icons-drawio generate-icons-drawio-dockerized verify test test-dockerized
