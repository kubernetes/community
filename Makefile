IMAGE_NAME=kube-communitydocs

default: \
	generate \

reset-docs:
	git checkout HEAD -- sig-list.md sig-*/README.md

build-image:
	docker build -q -t $(IMAGE_NAME) -f generator/Dockerfile generator

generate: build-image
	docker run --rm -e WHAT -v $(shell pwd):/go/src/app/generated:Z $(IMAGE_NAME) app

verify:
	@hack/verify.sh

test: build-image
	docker run --rm $(IMAGE_NAME) go test -v ./...

.PHONY: default reset-docs build-image generate verify test
