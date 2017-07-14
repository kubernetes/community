IMAGE_NAME=kube-communitydocs

all: \
	build-image \
	gen-docs \

reset-docs:
	git checkout HEAD -- sig-list.md sig-*

build-image:
	docker build -t $(IMAGE_NAME) -f generator/Dockerfile generator

gen-docs:
	docker run --rm -e WG -e SIG -v $(shell pwd):/go/src/app/generated $(IMAGE_NAME) app

test: build-image
	docker run --rm $(IMAGE_NAME) go test -v ./...
