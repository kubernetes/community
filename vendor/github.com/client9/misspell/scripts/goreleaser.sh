#!/bin/sh -e
# autorelease based on tag
test -n "$TRAVIS_TAG" && curl -sL https://git.io/goreleaser | bash
