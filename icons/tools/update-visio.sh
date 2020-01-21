#!/bin/bash

ICONS_PATH=${PWD}
DOCKER_IMAGE_REPO="k8s.gcr.io/visio-stencil-creator:v1.0"

# Run docker image that generates Visio stencil
# using png images.
docker run \
    -v ${ICONS_PATH}:/app/content \
    ${DOCKER_IMAGE_REPO} \
    "--image-pattern=**/labeled/*-256.png" \
    "--image-path=/app/content/png" \
    "--output-filename=/app/content/visio/kubernetes-visio-stencil.vssx"
