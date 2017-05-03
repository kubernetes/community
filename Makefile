all: \
	build-sigdocs \
	run-sigdocs \

reset-docs:
	git checkout HEAD -- sig-list.md sig-*

build-sigdocs:
	docker build -t sigdocs -f build/Dockerfile build

run-sigdocs:
	docker run -v $(shell pwd):/go/src/app sigdocs
