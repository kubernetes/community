all: \
	build-image \
	gen-docs \

reset-docs:
	git checkout HEAD -- sig-list.md sig-*

build-image:
	docker build -t sigdocs -f generator/Dockerfile generator

gen-doc:
	docker run -e WG=${WG} -e SIG=${SIG} -v $(shell pwd):/go/src/app sigdocs

gen-docs:
	docker run -v $(shell pwd):/go/src/app sigdocs

test:
	docker build -t sigdocs-test -f generator/Dockerfile.test generator
	docker run sigdocs-test
