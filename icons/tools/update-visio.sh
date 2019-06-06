#!/bin/bash

ICONS_PATH=${PWD}

GITHUB_ORG="hoveytechllc"
REPO_NAME="visio-stencil-creator"

rm -fdr ./tools/${REPO_NAME}

# Clone repository in current path
git clone https://github.com/${GITHUB_ORG}/${REPO_NAME}.git ./tools/${REPO_NAME}

# build image using Dockerfile from github repository
docker build \
    -t ${REPO_NAME}:latest \
    -f ./tools/${REPO_NAME}/Dockerfile \
    ./tools/${REPO_NAME}

# Run newly created Docker image
docker run \
    -v ${ICONS_PATH}:/app/content \
    ${REPO_NAME}:latest \
    "--image-pattern=**/labeled/*-256.png" \
    "--image-path=/app/content/png" \
    "--output-filename=/app/content/visio/kubernetes-visio-stencil.vssx"

rm -fdr ./tools/${REPO_NAME}